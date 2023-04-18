package Commands

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"

	BlockStructs "go-blockchain/blockstructs"
)

func ConfirmNFTTransaction(b *BlockStructs.Blockchain) {
	// Get the seller's private key
	var sellerPrivateKeyStr string
	fmt.Print("Enter your private key: ")
	fmt.Scan(&sellerPrivateKeyStr)

	// Get the buyer's public key
	var buyerPublicKeyStr string
	fmt.Print("Enter the buyer's public key: ")
	fmt.Scan(&buyerPublicKeyStr)

	// Get the NFT ID
	var nftID uint64
	fmt.Print("Enter the NFT ID: ")
	fmt.Scan(&nftID)

	// Get the transaction amount
	var amount float64
	fmt.Print("Enter the amount: ")
	fmt.Scan(&amount)

	// Find the pending NFT transaction
	nftTransaction, err := findPendingNFTTransaction(b, nftID, buyerPublicKeyStr, amount)
	if err != nil {
		fmt.Println("Error finding pending NFT transaction:", err)
		return
	}

	// Confirm the NFT transaction
	sellerPrivateKey, err := base64.StdEncoding.DecodeString(sellerPrivateKeyStr)
	if err != nil {
		fmt.Println("Error decoding private key:", err)
		return
	}

	// Sign the NFT transaction
	signature, err := signNFTTransaction(sellerPrivateKey, nftTransaction)
	if err != nil {
		fmt.Println("Error signing NFT transaction:", err)
		return
	}

	// Update the NFT transaction fields
	nftTransaction.Signature = signature
	nftTransaction.Confirmed = true


	// Create a new block with the pending NFT transactions
	b.NewNFTBlock(b.PendingNFTTransactions)

	// Clear the pending NFT transactions
	b.PendingNFTTransactions = []*BlockStructs.NFTTransaction{}

	fmt.Println("_______________________________________________________________")
	fmt.Println("NFT transaction confirmed, signed, and added to the blockchain")
	fmt.Println("_______________________________________________________________")
}

func findPendingNFTTransaction(b *BlockStructs.Blockchain, nftID uint64, buyerPublicKeyStr string, amount float64) (*BlockStructs.NFTTransaction, error) {
	for _, tx := range b.PendingNFTTransactions {
		if tx.NFTID == nftID && tx.ReceiverPubKey == buyerPublicKeyStr && tx.Amount == amount {
			return tx, nil
		}
	}

	return nil, errors.New("pending NFT transaction not found")
}

func signNFTTransaction(sellerPrivateKey []byte, nftTransaction *BlockStructs.NFTTransaction) (string, error) {
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
