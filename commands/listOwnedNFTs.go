package Commands

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/base64"
	"fmt"
	BlockStructs "go-blockchain/blockstructs"
	"math/big"
)

func ListOwnedNFTs(b *BlockStructs.Blockchain, privateKeyStr string) ([]*BlockStructs.CDKeyNFT, error) {
	// Decode the private key
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyStr)
	if err != nil {
		return nil, fmt.Errorf("error decoding private key: %v", err)
	}

	// Derive the public key from the private key
	curve := elliptic.P256()
	privateKey := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
		},
		D: new(big.Int).SetBytes(privateKeyBytes),
	}
	privateKey.PublicKey.X, privateKey.PublicKey.Y = curve.ScalarBaseMult(privateKey.D.Bytes())

	// Encode the public key to a string
	publicKeyBytes := elliptic.Marshal(curve, privateKey.PublicKey.X, privateKey.PublicKey.Y)
	publicKeyStr := base64.StdEncoding.EncodeToString(publicKeyBytes)

	// Find all NFTs owned by the wallet
	ownedNFTs := []*BlockStructs.CDKeyNFT{}
	for _, nft := range b.NFTs {
		if nft.OwnerPubKey == publicKeyStr {
			ownedNFTs = append(ownedNFTs, nft)
		}
	}

	return ownedNFTs, nil
}
