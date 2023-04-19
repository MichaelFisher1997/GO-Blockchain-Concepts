package BlockStructs

import (
	"encoding/base64"
	"fmt"
)

func Start() *Blockchain{
	// Create a new blockchain
	blockchain := NewBlockchain()

	// Create a new wallet
	var number float64
	fmt.Print("Enter Wallet amount: ")
			_, err := fmt.Scanf("%f", &number)
			if err != nil {
				fmt.Println("Error reading input:", err)
			}
	wallet := NewWallet(number)
	publicKeyStr := base64.StdEncoding.EncodeToString(wallet.PublicKey)

	fmt.Printf("New wallet created with private key: %v\n", base64.StdEncoding.EncodeToString(wallet.PrivateKey))
	fmt.Printf("New wallet created with public key: %v\n", base64.StdEncoding.EncodeToString(wallet.PublicKey))
	// Add the wallet to the blockchain
	blockchain.Wallets = append(blockchain.Wallets, wallet)
	// Add the wallet's public key to the list of authorities
	blockchain.Authorities = append(blockchain.Authorities, publicKeyStr)

	// Add the genesis block to the blockchain
	blockchain.NewGenesisBlock(publicKeyStr)
	// Display the latest block
	latestBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	fmt.Printf("Latest block: %v\n", latestBlock)

	// Check if the latest block has at least one transaction
	if len(latestBlock.Transactions) > 0 {
		fmt.Printf("Transaction in the latest block: %v\n", publicKeyStr)
	} else {
		fmt.Println("No transactions in the latest block.")
	}
	return blockchain
}