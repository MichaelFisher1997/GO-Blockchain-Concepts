package BlockStructs

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
)

func SignNFTTransaction(sellerPrivateKey []byte, nftTransaction *NFTTransaction) (string, error) {
	// Create an ECDSA private key from the provided bytes
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", err
	}

	// Create a hash of the NFT transaction
	txHash := nftTransaction.Hash()

	// Sign the transaction hash
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, txHash[:])
	if err != nil {
		return "", err
	}

	// Encode the signature to a base64 string
	signature := append(r.Bytes(), s.Bytes()...)
	signatureStr := base64.StdEncoding.EncodeToString(signature)

	return signatureStr, nil
}
