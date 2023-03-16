package BlockStructs

import (
	"fmt"
	//"os"
)

func Start() Blockchain{
	senderWallet := NewWallet(100.0) // The sender wallet will start with a balance of 100
	receiverWallet := NewWallet(50.0) // The receiver wallet will start with a balance of 0
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
	blockchain.UpdateBalances()  // Update the wallet balances
	return *blockchain

}