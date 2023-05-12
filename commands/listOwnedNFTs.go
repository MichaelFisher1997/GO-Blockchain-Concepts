package Commands

import (
	"fmt"
	BlockStructs "go-blockchain/blockstructs"
	Utils "go-blockchain/utils"
)

func ListOwnedNFTs(b *BlockStructs.Blockchain, privateKeyStr string) ([]*BlockStructs.CDKeyNFT) {
	publicKeyStr, err := Utils.DecodePrivateKey(privateKeyStr)
	Utils.Check(err)
	// Find all NFTs owned by the wallet
	ownedNFTs := []*BlockStructs.CDKeyNFT{}
	for _, nft := range b.NFTs {
		fmt.Println("nft.OwnerPubKey: ", nft.OwnerPubKey)
		if nft.OwnerPubKey == publicKeyStr {
			ownedNFTs = append(ownedNFTs, nft)
		}
	}

	return ownedNFTs
}
