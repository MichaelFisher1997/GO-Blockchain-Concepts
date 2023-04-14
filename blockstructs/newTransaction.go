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

/*func NewTransaction(senderPrivateKey *ecdsa.PrivateKey, recipientPublicKeyStr string, amount float64) *Transaction {
	// Decode the recipient's public key
	recipientPublicKeyBytes, err := base64.StdEncoding.DecodeString(recipientPublicKeyStr)
	if err != nil {
		fmt.Println("Error decoding recipient public key:", err)
		return nil
	}

	// Convert the decoded public key bytes to an ecdsa.PublicKey
	curve := elliptic.P256()
	x, y := elliptic.Unmarshal(curve, recipientPublicKeyBytes)
	recipientPublicKey := &ecdsa.PublicKey{Curve: curve, X: x, Y: y}

	// Create the transaction
	tx := &Transaction{
		SenderPublicKey:    senderPrivateKey.PublicKey,
		ReceiverPublicKey:  *recipientPublicKey,
		Amount:             amount,
	}

	// Sign the transaction
	tx.Sign(senderPrivateKey)

	return tx
}*/
