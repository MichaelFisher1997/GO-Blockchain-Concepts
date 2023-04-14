package Commands

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/base64"
	"fmt"
	BlockStructs "go-blockchain/blockstructs"
)

func NewCDKeyNFT(id uint64, cdKey string, tokenID uint64, ownerPublicKey []byte) (*BlockStructs.CDKeyNFT, error) {
	// Get the elliptic curve
	curve := elliptic.P256()

	// Unmarshal the public key
	x, y := elliptic.Unmarshal(curve, ownerPublicKey)
	if x == nil || y == nil {
		return nil, fmt.Errorf("failed to unmarshal the public key")
	}

	// Create an ecdsa.PublicKey
	ownerWallet := &ecdsa.PublicKey{
		Curve: curve,
		X:     x,
		Y:     y,
	}

	// Encode the public key to a base64 string
	publicKeyBytes := elliptic.Marshal(curve, ownerWallet.X, ownerWallet.Y)
	publicKeyStr := base64.StdEncoding.EncodeToString(publicKeyBytes)

	// Create a new CDKeyNFT
	nft := &BlockStructs.CDKeyNFT{
		ID:          id,
		CDKey:       cdKey,
		TokenID:     tokenID,
		Minted:      true,
		MintedBy:    publicKeyStr, // Use the base64-encoded public key string
		MintedOn:    BlockStructs.TimeStamp(),
		OwnerPubKey: publicKeyStr,
	}

	return nft, nil
}

/*
input number <- : 5
Enter the ID: 123
Enter the CD Key: AB12-CD34-EF56-GH78
Enter the Token ID: 1001
Enter your private key: WgFXmDQZ8aHtGmRjI9X9trLmPp8WxjKq1cbrDtI7Npk=
*/
