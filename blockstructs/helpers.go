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