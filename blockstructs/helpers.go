package BlockStructs

import "bytes"

func (b *Blockchain) findWalletByPublicKey(publicKey []byte) *Wallet {
	for _, wallet := range b.Wallets {
		if bytes.Equal(wallet.PublicKey, publicKey) {
			return wallet
		}
	}
	return nil
}

func (b *Blockchain) findNFTByID(nftID uint64) *CDKeyNFT {
	for _, nft := range b.NFTs {
		if nft.ID == nftID {
			return nft
		}
	}
	return nil
}
