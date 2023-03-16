package BlockStructs

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"math/big"
)

// SignTransaction signs a transaction using the sender's private key.
// The private key is passed as a byte slice, and the signature is appended to the transaction.
func SignTransaction(transaction *Transaction, senderPrivateKey []byte) error {
    // Compute the hash of the transaction
    hash := transactionHash(transaction)

    // Create a new elliptic curve instance for the P-256 curve
    curve := elliptic.P256()

    // Compute the public key points (x, y) corresponding to the private key
    x, y := curve.ScalarBaseMult(senderPrivateKey)

    // Reconstruct the private key as an ecdsa.PrivateKey instance
    privateKey := ecdsa.PrivateKey{
        PublicKey: ecdsa.PublicKey{
            Curve: curve,
            X:     x,
            Y:     y,
        },
        D: new(big.Int).SetBytes(senderPrivateKey),
    }

    // Sign the transaction hash using the reconstructed private key
    r, s, err := ecdsa.Sign(rand.Reader, &privateKey, hash[:])
    if err != nil {
        return err
    }

    // Combine the signature components (r, s) into a single byte slice
    signature := append(r.Bytes(), s.Bytes()...)

    // Append the signature to the transaction
    transaction.Signature = signature
    return nil
}



// transactionHash calculates the SHA-256 hash of the transaction data.
// It returns the hash as an array of 32 bytes.
func transactionHash(transaction *Transaction) [32]byte {
	// Serialize the transaction data as JSON.
	txData, err := json.Marshal(transaction)
	if err != nil {
		panic(err)
	}

	// Calculate and return the SHA-256 hash of the serialized transaction data.
	return sha256.Sum256(txData)
}

func IsValidTransaction(transaction *Transaction) bool {
	// Deserialize the public key from the byte slice
	curve := elliptic.P256()
	x, y := elliptic.Unmarshal(curve, transaction.SenderPublicKey)
	senderPublicKey := ecdsa.PublicKey{
		Curve: curve,
		X:     x,
		Y:     y,
	}

	// Check if the sender and receiver are not the same
	if bytes.Equal(transaction.SenderPublicKey, transaction.ReceiverPublicKey) {
		return false
	}

	// Verify the transaction signature
	hash := transactionHash(transaction)
	r := new(big.Int).SetBytes(transaction.Signature[:len(transaction.Signature)/2])
	s := new(big.Int).SetBytes(transaction.Signature[len(transaction.Signature)/2:])
	return ecdsa.Verify(&senderPublicKey, hash[:], r, s)
}






