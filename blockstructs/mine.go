package BlockStructs

import (
	"encoding/base64"
	"errors"
	"fmt"
)

func (b *Blockchain) Mine(){
	b.ProcessNewNft()
	err := b.ProcessPendingNFTTransactions()
	fmt.Println("err:",err)
	fmt.Println("MINE! :)")

	
} 

func (b  *Blockchain) ProcessNewNft(){ 
	// Check if there are pending NFT transactions
	if b.PendingNFTs == nil {
		b.NFTs = []*CDKeyNFT{} //make sure to clear
		return 
	} else {
		for _, NFTs := range b.PendingNFTs {
			// Check if the creator of the NFT is the owner of the NFT
			if NFTs.MintedBy == NFTs.OwnerPubKey && b.FindNFTByID(NFTs.ID) == nil{
				// Append the transaction to the list of pending transactions
				NFTs.Minted = true
				b.NFTs = append(b.NFTs, NFTs)
			}
			
		}
		// Clear the pending NFT transactions
		b.PendingNFTs = []*CDKeyNFT{}
		
		return 
	}
}
func (b *Blockchain) ProcessPendingNFTTransactions() error {
	if len(b.PendingNFTTransactions) == 0 {
		b.PendingNFTTransactions = []*NFTTransaction{}
		return nil
	}

	for _, tx := range b.PendingNFTTransactions {
		nft := b.FindNFTByID(tx.NFTID)

		if nft == nil {
			return errors.New("NFT not found")
		}

		err := MakeTransactionWithDetails2(b, tx.OwnerPubKey, tx.ReceiverPubKey, tx.Amount)
		if err != nil {
			fmt.Println("Error making transaction:", err)
		}

		// Update the owner of the NFT
		nft.OwnerPubKey = tx.ReceiverPubKey
	}

	// Clear the pending NFT transactions
	b.PendingNFTTransactions = []*NFTTransaction{}

	return nil
}

func MakeTransactionWithDetails2(b *Blockchain, senderPublicKey string, recipientPublicKeyStr string, amount float64) error {
	senderPublicKeyBytes, err := base64.StdEncoding.DecodeString(senderPublicKey)
	if err != nil {
		return errors.New("Error decoding sender's public key")
	}
	senderWallet := b.FindWalletByPublicKey(senderPublicKeyBytes)

	recipientPublicKeyBytes, err := base64.StdEncoding.DecodeString(recipientPublicKeyStr)
	if err != nil {
		return errors.New("Error decoding recipient's public key")
	}
	recipientWallet := b.FindWalletByPublicKey(recipientPublicKeyBytes)


	// Create the transaction
	transaction := NewTransaction(senderWallet, recipientWallet, amount)
	b.PendingTransactions = append(b.PendingTransactions, transaction)

	ValidateAndAddTransactions(b)

	return nil
}










