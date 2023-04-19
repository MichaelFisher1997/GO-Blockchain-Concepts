package BlockStructs

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/base64"
	"fmt"
	Utils "go-blockchain/utils"
	"math/big"
)

// ...

func AddToAuthorities(b *Blockchain, privateKeyStr string) {
	if !b.ValidateBlockchain() {
        fmt.Println("Error: The blockchain is not valid.")
        return
    }
	// Decode the private key and derive the public key
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyStr)
	Utils.Check(err)
	curve := elliptic.P256()
	privateKey := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
		},
		D: new(big.Int).SetBytes(privateKeyBytes),
	}
	privateKey.PublicKey.X, privateKey.PublicKey.Y = curve.ScalarBaseMult(privateKey.D.Bytes())
	publicKeyBytes := elliptic.Marshal(curve, privateKey.PublicKey.X, privateKey.PublicKey.Y)
	publicKeyStr := base64.StdEncoding.EncodeToString(publicKeyBytes)

	// Check if the public key is already in the Authorities list
	for _, authority := range b.Authorities {
		if authority == publicKeyStr {
			fmt.Println("Public key is already in the Authorities list.")
			return
		}
	}

	// Add the public key to the Authorities list
	b.Authorities = append(b.Authorities, publicKeyStr)
	fmt.Println("Public key has been added to the Authorities list.")
}

func (b *Blockchain) ValidateBlockchain() bool {
    // Check if the blockchain is empty
    if len(b.Blocks) == 0 {
        fmt.Println("Error: The blockchain is empty.")
        return false
    }

    // Validate each block in the blockchain
    for i := 1; i < len(b.Blocks); i++ {
        currentBlock := b.Blocks[i]
        previousBlock := b.Blocks[i-1]

        // Check if the current block's HashPrevBlock is equal to the previous block's hash
        if currentBlock.HashPrevBlock != previousBlock.BlockHash() {
            fmt.Printf("Error: Block %d has an invalid HashPrevBlock value.\n", currentBlock.BlockID)
            return false
        }
    }

    return true
}


