package Utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"encoding/base64"
	"math/big"
)

func DecodePrivateKey(privateKeyStr string) (publicKeyStr string, err error) {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyStr)
	if err != nil {
		return "", err
	}

	curve := elliptic.P256()
	privateKey := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
		},
		D: new(big.Int).SetBytes(privateKeyBytes),
	}
	privateKey.PublicKey.X, privateKey.PublicKey.Y = curve.ScalarBaseMult(privateKey.D.Bytes())
	publicKeyBytes := elliptic.Marshal(curve, privateKey.PublicKey.X, privateKey.PublicKey.Y)
	publicKeyStr = base64.StdEncoding.EncodeToString(publicKeyBytes)

	return publicKeyStr, nil
}

// VerifySignature verifies a signature given a public key, signature, and data hash
func VerifySignature(publicKeyStr, signatureStr string, dataHash [32]byte) bool {
	// Decode the public key
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyStr)
	if err != nil {
		return false
	}

	publicKeyInterface, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		return false
	}

	publicKey, ok := publicKeyInterface.(*ecdsa.PublicKey)
	if !ok {
		return false
	}

	// Decode the signature
	signatureBytes, err := base64.StdEncoding.DecodeString(signatureStr)
	if err != nil {
		return false
	}

	r := new(big.Int).SetBytes(signatureBytes[:len(signatureBytes)/2])
	s := new(big.Int).SetBytes(signatureBytes[len(signatureBytes)/2:])

	// Verify the signature
	return ecdsa.Verify(publicKey, dataHash[:], r, s)
}
