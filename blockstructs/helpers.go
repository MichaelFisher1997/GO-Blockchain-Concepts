package BlockStructs

import (
	"bytes"
	//Commands "go-blockchain/commands"
)

func (b *Blockchain) FindWalletByPublicKey(publicKey []byte) *Wallet {
	for _, wallet := range b.Wallets {
		if bytes.Equal(wallet.PublicKey, publicKey) {
			return wallet
		}
	}
	return nil
}

func (b *Blockchain) FindNFTByID(nftID uint64) *CDKeyNFT {
	for _, nft := range b.NFTs {
		if nft.ID == nftID {
			return nft
		}
	}
	return nil
}

/*func ProcessPendingNFTs(b *Blockchain, creatorPubKey string) {
    for {
        time.Sleep(1 * time.Second)

        if len(b.PendingNFTs) > 0 {
            // Create a new block with the pending NFTs
            err := NewNFTBlock(b.PendingNFTs, creatorPubKey)
            if err != nil {
                log.Printf("Error creating NFT block: %v", err)
            } else {
                // Clear the pending NFTs after creating the new block
                b.PendingNFTs = []*CDKeyNFT{}
            }
        }
    }
}*/
