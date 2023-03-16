package BlockStructs

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"errors"
)

func (blockchain *Blockchain) GetBalance(publicKey []byte) float64 {
	balance := 0.0

	// Deserialize the target public key from the byte slice
	curve := elliptic.P256()
	x, y := elliptic.Unmarshal(curve, publicKey)
	targetPublicKey := ecdsa.PublicKey{
		Curve: curve,
		X:     x,
		Y:     y,
	}

	// Iterate through the blocks in the blockchain
	for _, block := range blockchain.Blocks {
		// Iterate through the transactions in each block
		for _, tx := range block.Transactions {
			// Deserialize the sender public key from the byte slice
			senderX, senderY := elliptic.Unmarshal(curve, tx.SenderPublicKey)
			senderPublicKey := ecdsa.PublicKey{
				Curve: curve,
				X:     senderX,
				Y:     senderY,
			}

			// Deserialize the receiver public key from the byte slice
			receiverX, receiverY := elliptic.Unmarshal(curve, tx.ReceiverPublicKey)
			receiverPublicKey := ecdsa.PublicKey{
				Curve: curve,
				X:     receiverX,
				Y:     receiverY,
			}

			// If the public key is the sender, subtract the transaction amount from the balance
			if targetPublicKey == senderPublicKey {
				balance -= tx.Amount
			}

			// If the public key is the receiver, add the transaction amount to the balance
			if targetPublicKey == receiverPublicKey {
				balance += tx.Amount
			}
		}
	}

	return balance
}

// GetWallet finds a wallet in the blockchain's wallets slice using the given public key.
// It returns the wallet if found, and an error if not found.
func (b *Blockchain) GetWallet(publicKey []byte) (*Wallet, error) {
    // Iterate through the wallets in the blockchain
    for _, wallet := range b.Wallets {
        // If the wallet's public key matches the given public key, return the wallet
        if bytes.Equal(wallet.PublicKey, publicKey) {
            return wallet, nil
        }
    }
    // If no matching wallet is found, return an error
    return nil, errors.New("wallet not found")
}



func (b *Blockchain) UpdateBalances() {
	for _, block := range b.Blocks {
		for _, tx := range block.Transactions {
			// Decrease the sender's balance
			senderWallet, _ := b.GetWallet(tx.SenderPublicKey)
			if senderWallet != nil {
				senderWallet.Balance -= tx.Amount
			}

			// Increase the receiver's balance
			receiverWallet, _ := b.GetWallet(tx.ReceiverPublicKey)
			if receiverWallet != nil {
				receiverWallet.Balance += tx.Amount
			}
		}
	}
}


