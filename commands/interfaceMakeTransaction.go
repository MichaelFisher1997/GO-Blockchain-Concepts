package Commands

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/base64"
	"fmt"
	BlockStructs "go-blockchain/blockstructs"
	"math/big"
)

func MakeTransaction(b *BlockStructs.Blockchain) {
	// Get the private key from the user
	var privateKeyStr string
	fmt.Print("Enter your private key: ")
	fmt.Scan(&privateKeyStr)

	// Decode the private key
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyStr)
	if err != nil {
		fmt.Println("Error decoding private key:", err)
		return
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

	// Get the recipient's public key from the user
	var recipientPublicKeyStr string
	fmt.Print("Enter the recipient's public key: ")
	fmt.Scan(&recipientPublicKeyStr)

	//Get the amount to send from the user
	var amount float64
	fmt.Print("Enter the amount to send: ")
	_, err = fmt.Scanf("%f", &amount)
	if err != nil {
		fmt.Println("Error reading input:", err)
	}

	// Decode the recipient's public key
	recipientPublicKeyBytes, err := base64.StdEncoding.DecodeString(recipientPublicKeyStr)
	if err != nil {
		fmt.Println("Error decoding recipient's public key:", err)
		return
	}

	recipientWallet := &BlockStructs.Wallet{
		PublicKey: recipientPublicKeyBytes,
	}

	// Create the transaction
	transaction := BlockStructs.NewTransaction(senderWallet, recipientWallet, amount)
	b.PendingTransactions = append(b.PendingTransactions, transaction)

	BlockStructs.ValidateAndAddTransactions(b)
}
