package Commands

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/base64"
	"fmt"
	BlockStructs "go-blockchain/blockstructs"
	"math/big"
)

func GetBalanceByPrivateKey(b *BlockStructs.Blockchain, privateKeyStr string) (float64, error) {
	// Decode the private key
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyStr)
	if err != nil {
		return 0, fmt.Errorf("Error decoding private key: %v", err)
	}

	// Derive the public key from the private key
	curve := elliptic.P256()
	privateKey := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
		},
		D: new(big.Int).SetBytes(privateKeyBytes),
	}
	privateKey.PublicKey.X, privateKey.PublicKey.Y = curve.ScalarBaseMult(privateKey.D.Bytes())

	// Encode the public key to a string
	publicKeyBytes := elliptic.Marshal(curve, privateKey.PublicKey.X, privateKey.PublicKey.Y)
	publicKeyStr := base64.StdEncoding.EncodeToString(publicKeyBytes)

	// Find the corresponding wallet in the blockchain
	var wallet *BlockStructs.Wallet
	for _, w := range b.Wallets {
		if publicKeyStr == base64.StdEncoding.EncodeToString(w.PublicKey) {
			wallet = w
			break
		}
	}

	if wallet == nil {
		return 0, fmt.Errorf("Wallet not found")
	}

	return wallet.Balance, nil
}
