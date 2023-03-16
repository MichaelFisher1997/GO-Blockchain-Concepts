package BlockStructs

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

// NewWallet creates a new wallet with a public and private key pair.
func NewWallet(initialBalance float64) *Wallet {
    privateKey, publicKey := generateKeyPair()
    publicKeyBytes := elliptic.Marshal(publicKey.Curve, publicKey.X, publicKey.Y)
    return &Wallet{PublicKey: publicKeyBytes, PrivateKey: privateKey.D.Bytes(), Balance: initialBalance}
}



// generateKeyPair generates a new public and private key pair using the ECDSA
// algorithm with the P-256 curve. It returns the private and public keys.
func generateKeyPair() (ecdsa.PrivateKey, ecdsa.PublicKey) {
	// Use the P-256 elliptic curve for the key pair generation.
	curve := elliptic.P256()

	// Generate a new private key using the P-256 curve and the cryptographically
	// secure random number generator.
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	
	// Derive the public key from the private key.
	publicKey := privateKey.PublicKey

	// Return the generated private and public keys.
	return *privateKey, publicKey
}
