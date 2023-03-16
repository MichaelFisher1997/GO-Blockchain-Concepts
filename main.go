package main

import (
	"fmt"
	BlockStructs "go-blockchain/blockstructs"

	Read "go-blockchain/read"
	Utils "go-blockchain/utils"
	//"os"
)

func main() {
	fmt.Println("------start------")

	senderWallet := BlockStructs.NewWallet(100.0) // The sender wallet will start with a balance of 100
	receiverWallet := BlockStructs.NewWallet(50.0) // The receiver wallet will start with a balance of 0

	// Attempt to load the blockchain from the binary file.
	blockchain, err := Read.LoadBlockchain("blockchain.dat")
	if err != nil {
		// If there is an error (e.g., the file doesn't exist), create a new blockchain with the sender and receiver wallets.
        blockchain = BlockStructs.NewBlockchain()
	}
	// Create two new wallets

	// Add wallets to the blockchain
	blockchain.Wallets = append(blockchain.Wallets, senderWallet, receiverWallet)
	//

	// Create a new transaction from the sender to the receiver
	transaction := &BlockStructs.Transaction{
		SenderPublicKey:    senderWallet.PublicKey,
		ReceiverPublicKey:  receiverWallet.PublicKey,
		Amount:             10.0, // Set the transaction amount
	}

	// Sign the transaction using the sender's private key
	err2 := BlockStructs.SignTransaction(transaction, senderWallet.PrivateKey)

	if err2 != nil {
		fmt.Printf("Error signing transaction: %v\n", err2)
		return
	}

	// Add the transaction to the blockchain
	transactions := []*BlockStructs.Transaction{transaction}

	// Create a new block with the transaction
    blockchain.NewBlock(transactions)
	blockchain.UpdateBalances()  // Update the wallet balances

	// Print the latest block in the blockchain
	latestBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	fmt.Printf("Latest block: %v\n", latestBlock)

	// Print the transaction in the latest block
	fmt.Printf("Transaction in the latest block: %v\n", latestBlock.Transactions[0])

	//----------------SAVE-------------
	//time.Sleep(1 * time.Second)
	err4 := Read.SaveToFile("ledger/myBlocks.json",blockchain)
	Utils.Check(err4)
	// Save the current state of the blockchain to a binary file.
	if err := Read.SaveBlockchain(blockchain, "blockchain.dat"); err != nil {
		fmt.Printf("Error saving blockchain: %v\n", err)
	}
}

