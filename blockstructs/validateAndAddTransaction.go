package BlockStructs

import (
	"fmt"
)

func ValidateAndAddTransactions(b *Blockchain) {
	validTransactions := []*Transaction{}

	for _, transaction := range b.PendingTransactions {
		senderPublicKey := transaction.SenderPublicKey
		recipientPublicKey := transaction.ReceiverPublicKey
		amount := transaction.Amount

		// Find the sender's wallet in the blockchain
		var senderWallet *Wallet
		for _, wallet := range b.Wallets {
			if string(wallet.PublicKey) == string(senderPublicKey) {
				senderWallet = wallet
				break
			}
		}

		if senderWallet == nil {
			fmt.Println("Sender wallet not found")
			continue
		}

		// Validate the transaction
		if senderWallet.Balance >= amount {
			// Add the transaction to the valid transactions list
			validTransactions = append(validTransactions, transaction)

			// Update the sender's balance
			senderWallet.Balance -= amount

			// Find the recipient's wallet in the blockchain
			var recipientWallet *Wallet
			for _, wallet := range b.Wallets {
				if string(wallet.PublicKey) == string(recipientPublicKey) {
					recipientWallet = wallet
					break
				}
			}

			// If the recipient's wallet is not found, create a new one
			if recipientWallet == nil {
				recipientWallet = &Wallet{
					PublicKey: recipientPublicKey,
					Balance:   0,
				}
				b.Wallets = append(b.Wallets, recipientWallet)
			}

			// Update the recipient's balance
			recipientWallet.Balance += amount
		} else {
			fmt.Println("____________________________________________________")
			fmt.Println("Insufficient funds for the transaction")
			fmt.Println("____________________________________________________")
		}
	}

	// Create a new block with valid transactions
	b.NewBlock(validTransactions)

	// Clear the pending transactions
	b.PendingTransactions = []*Transaction{}
}
