package BlockStructs

import (
	"fmt"
	//"os"
)

func Start() *Blockchain{
	// Create a new blockchain
	blockchain := NewBlockchain()
	latestBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	fmt.Printf("Latest block: %v\n", latestBlock)

	// Check if the latest block has at least one transaction
	if len(latestBlock.Transactions) > 0 {
		fmt.Printf("Transaction in the latest block: %v\n", latestBlock.Transactions[0])
	} else {
		fmt.Println("No transactions in the latest block.")
	}
	return blockchain

}

/*
senderWallet := NewWallet(100.0) // The sender wallet will start with a balance of 100
	receiverWallet := NewWallet(0.0) // The receiver wallet will start with a balance of 0
	blockchain := NewBlockchain()
	// Add wallets to the blockchain
	blockchain.Wallets = append(blockchain.Wallets, senderWallet, receiverWallet)
	// Create a new transaction from the sender to the receiver
	transaction := &Transaction{
		SenderPublicKey:    senderWallet.PublicKey,
		ReceiverPublicKey:  receiverWallet.PublicKey,
		Amount:             10.0, // Set the transaction amount
	}
	// Sign the transaction using the sender's private key
	err2 := SignTransaction(transaction, senderWallet.PrivateKey)
	if err2 != nil {
		fmt.Printf("Error signing transaction: %v\n", err2)
	}
	// Add the transaction to the blockchain
	transactions := []*Transaction{transaction}
	// Create a new block with the transaction
    blockchain.NewBlock(transactions)
	//blockchain.UpdateBalances()  // Update the wallet balances
	latestBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	fmt.Printf("Latest block: %v\n", latestBlock)
	fmt.Printf("Transaction in the latest block: %v\n", latestBlock.Transactions[0])
	return blockchain
	*/