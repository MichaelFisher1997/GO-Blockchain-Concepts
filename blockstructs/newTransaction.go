package BlockStructs

import (
	"fmt"
)


func NewTransaction(senderWallet *Wallet, receiverWallet *Wallet, amount float64) *Transaction {
	// Create a new transaction from the sender to the receiver
	transaction := &Transaction{
		SenderPublicKey:    senderWallet.PublicKey,
		ReceiverPublicKey:  receiverWallet.PublicKey,
		Amount:             amount, // Set the transaction amount
	}
	// Sign the transaction using the sender's private key
	err2 := SignTransaction(transaction, senderWallet.PrivateKey)
	if err2 != nil {
		fmt.Printf("Error signing transaction: %v\n", err2)
	}
	return transaction
}