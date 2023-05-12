package BlockStructs

import (
	"fmt"
	Utils "go-blockchain/utils"
)

func (b *Blockchain)PendingNFTTransaction(PrivateKey string, NFT_ID uint64, Price float64) (error) {
	// Seed the random number generator with the current time

	//get the public key
	PublicKey, _ := Utils.DecodePrivateKey(PrivateKey)
	//create the NFT transaction
	for _, NFT := range b.NFTs {
		if NFT.ID == NFT_ID {
			NFTTransaction := &NFTTransaction{
					ID:             uint64(Utils.RandomNumber()),
					NFTID:          NFT_ID,
					OwnerPubKey:   PublicKey,
					ReceiverPubKey: "",
					Amount:         Price,
					Signature:      "",
					Confirmed:      false,
					TimeStamp:      TimeStamp(),
				}

			// Append the transaction to the list of pending transactions
			b.PendingNFTTransactions = append(b.PendingNFTTransactions, NFTTransaction)

			return nil
		}
		
		
		
	}
	err := fmt.Errorf("NFT not found")
	return err
}

// FindPendingNFTTransaction finds a pending NFT transaction in the blockchain
func (b *Blockchain) AllPendingNFTTransactions() []*NFT_Adds {
	adds := []*NFT_Adds{}
	for _, transaction := range b.PendingNFTTransactions {
		for _, nft := range b.NFTs {
			if nft.ID == transaction.NFTID {
				nftAdd := &NFT_Adds{
					ID:      transaction.ID,
					TokenID: nft.CDKey,
					Price:   transaction.Amount,
					Time:    TimeStamp(), // assuming your NFTTransaction struct has a TimeStamp field
				}
				adds = append(adds, nftAdd)
				break // we found the matching NFT, no need to continue the inner loop
			}
		}
	}
	return adds
}

func (b *Blockchain) BuyNFT(nftID uint64, buyerPrivateKeyStr string, amount float64) error {
	// Decode the buyer's private key string
	buyerPrivateKey, err := Utils.DecodePrivateKey(buyerPrivateKeyStr)
	if err != nil {
		return err
	}

	// Find the pending NFT transaction
	found := false
	for _, tx := range b.PendingNFTTransactions {
		if tx.NFTID == nftID && tx.Amount == amount {
			tx.ReceiverPubKey = buyerPrivateKey
			signature, err := SignNFTTransaction([]byte(buyerPrivateKey), tx)
			if err != nil {
				return err
			}
			tx.Signature = signature
			found = true
			break
		}
	}
	// If the transaction was not found, return an error
	if !found {
		return fmt.Errorf("NFT transaction not found")
	}

	return nil
}





