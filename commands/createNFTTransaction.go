package Commands

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/base64"
	"fmt"
	BlockStructs "go-blockchain/blockstructs"
	"math/big"
)

func CreateNFTTransaction(b *BlockStructs.Blockchain) {
	// Get buyer's private key
	var buyerPrivateKeyStr string
	fmt.Print("Enter your private key: ")
	fmt.Scan(&buyerPrivateKeyStr)

	// Get NFT ID
	var nftID uint64
	fmt.Print("Enter the NFT ID: ")
	fmt.Scan(&nftID)

	// Get the seller's public key
	var sellerPublicKeyStr string
	fmt.Print("Enter the seller's public key: ")
	fmt.Scan(&sellerPublicKeyStr)

	// Get the transaction amount
	var amount float64
	fmt.Print("Enter the amount: ")
	fmt.Scan(&amount)

	// Derive the buyer's public key from the private key
	buyerPublicKey, err := publicKeyFromPrivateKey(buyerPrivateKeyStr)
	if err != nil {
		fmt.Println("Error deriving public key:", err)
		return
	}

	// Create the NFT transaction
	nftTransaction := &BlockStructs.NFTTransaction{
		ID:             0, // You may want to generate a unique ID for the transaction
		NFTID:          nftID,
		SenderPubKey:   sellerPublicKeyStr,
		ReceiverPubKey: buyerPublicKey,
		Amount:         amount,
		Signature:      "", // The buyer can sign the transaction, but it's not mandatory at this stage
		Confirmed:      false,
	}

	// Add the NFT transaction to the PendingNFTTransactions list
	b.PendingNFTTransactions = append(b.PendingNFTTransactions, nftTransaction)

	fmt.Println("NFT transaction created and added to pending transactions")
}

func publicKeyFromPrivateKey(privateKeyStr string) (string, error) {
	// Decode the private key
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyStr)
	if err != nil {
		return "", err
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

	return publicKeyStr, nil
}
