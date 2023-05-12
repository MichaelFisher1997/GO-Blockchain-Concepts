package Commands

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/base64"
	"errors"
	BlockStructs "go-blockchain/blockstructs"
	"math/big"
)

//API version
func MakeTransactionWithDetails(b *BlockStructs.Blockchain, privateKeyStr string, recipientPublicKeyStr string, amount float64) error {
	// Decode the private key
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyStr)
	if err != nil {
		return errors.New("Error decoding private key")
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

	// Encode the public key to a byte slice
	publicKeyBytes := elliptic.Marshal(curve, privateKey.PublicKey.X, privateKey.PublicKey.Y)

	// Derive the sender wallet from the private key
	senderWallet := &BlockStructs.Wallet{
		PublicKey:  publicKeyBytes,
		PrivateKey: privateKey.D.Bytes(),
	}

	// Decode the recipient's public key
	recipientPublicKeyBytes, err := base64.StdEncoding.DecodeString(recipientPublicKeyStr)
	if err != nil {
		return errors.New("Error decoding recipient's public key")
	}

	recipientWallet := &BlockStructs.Wallet{
		PublicKey: recipientPublicKeyBytes,
	}

	// Create the transaction
	transaction := BlockStructs.NewTransaction(senderWallet, recipientWallet, amount)
	b.PendingTransactions = append(b.PendingTransactions, transaction)

	BlockStructs.ValidateAndAddTransactions(b)

	return nil
}