package BlockStructs

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func (b *Blockchain) NewBlock(transactions []*Transaction, creatorPubKey string) {
	tempBlock := &Block{
        CreatorPubKey: creatorPubKey,
    }

    if !b.IsValidBlock(tempBlock) {
        fmt.Println("Error: Block creator is not an authorized authority.")
        return
    }
	b.BlockCount = b.BlockCount+1
	//PrevBlockHash :=  ""

	block := &Block{
		Magic_No: "0xF9S834SK",
		BlockID: b.BlockCount,
		Blocksize: 80,
		Version:		1,
		HashPrevBlock:  b.PrevBlockHash(),
		HashMerkleRoot: b.MerkelRoot(),
		TimeStamp:      TimeStamp(),
		Transaction_counter: len(transactions),
        Transactions:        transactions,
		CreatorPubKey: creatorPubKey,
	}
	fmt.Println("New block :",block.BlockHash())
	b.Blocks = append(b.Blocks, block)
	// Clear the pending transactions after adding them to the new block
	b.PendingTransactions = []*Transaction{}
}

func (b *Blockchain) NewNFTBlock(nftTransactions []*NFTTransaction, creatorPubKey string) {
	tempBlock := &Block{
        CreatorPubKey: creatorPubKey,
    }

    if !b.IsValidBlock(tempBlock) {
        fmt.Println("Error: Block creator is not an authorized authority.")
        return
    }

	b.BlockCount = b.BlockCount + 1

	block := &Block{
		Magic_No:          "0xF9S834SK",
		BlockID:           b.BlockCount,
		Blocksize:         80,
		Version:           1,
		HashPrevBlock:     b.PrevBlockHash(),
		HashMerkleRoot:    b.MerkelRoot(),
		TimeStamp:         TimeStamp(),
		Transaction_counter: len(nftTransactions),
		NFTTransactions:     nftTransactions,
		CreatorPubKey: creatorPubKey,
	}
	// Update wallet balances and NFT ownership
	for _, nftTransaction := range nftTransactions {
		if nftTransaction.Confirmed {
			// Deduct the amount from the seller's wallet
			senderPublicKeyBytes, err := base64.StdEncoding.DecodeString(nftTransaction.SenderPubKey)
			if err != nil {
				fmt.Println("Error decoding seller's public key:", err)
				return
			}
			sellerWallet := b.findWalletByPublicKey(senderPublicKeyBytes)
			if sellerWallet != nil {
				sellerWallet.Balance -= nftTransaction.Amount
			}

			// Decode the buyer's public key from base64
			receiverPublicKeyBytes, err := base64.StdEncoding.DecodeString(nftTransaction.ReceiverPubKey)
			if err != nil {
				fmt.Println("Error decoding buyer's public key:", err)
				return
			}

			// Add the amount to the buyer's wallet
			buyerWallet := b.findWalletByPublicKey(receiverPublicKeyBytes)
			if buyerWallet != nil {
				buyerWallet.Balance += nftTransaction.Amount
			}

			// Update the NFT's ownership
			nft := b.findNFTByID(nftTransaction.NFTID)
			if nft != nil {
				nft.OwnerPubKey = nftTransaction.ReceiverPubKey
			}
		}
	}
	fmt.Println("New NFT block :", block.BlockHash())
	b.Blocks = append(b.Blocks, block)
	// Clear the pending NFT transactions after adding them to the new block
	b.PendingNFTTransactions = []*NFTTransaction{}
}







func (b *Block) BlockHash() string {
	//convert to string
	header := fmt.Sprintf("%v" ,b)
	h := sha256.Sum256([]byte(header))
	return hex.EncodeToString(h[:])
}
/*
BCJGek3C6GFO2Wr9OFzI+sw6mkRfeoVlkv3357QWfbEsCMf4XM2f0kdR8gNxeW7BB9MwLwmpkuWUbEMNDqxLdwA=
BIVJjhFCRJnkSKE6w0Sli5GFv549HTka2kDRRXEg61yH/4XMekaoTyy3lc4gEuHY9e4Ef8dAISMSX+ylZbM2ikk=
*/
func NewBlockchain() *Blockchain {
    blockchain := &Blockchain{
        Wallets: []*Wallet{},
		Authorities: []string{},
    }

    // Add the genesis block to the blockchain
    //blockchain.NewGenesisBlock()

    return blockchain
}


func (b *Blockchain) NewGenesisBlock(creatorPubKey string) {
	block := &Block{
		Magic_No: "0xD9B4BEF9",
		BlockID: 0,
		Blocksize: 80,
		Version:        1, //You upgrade the software and it specifies a new version
		HashPrevBlock:  "0000000000000000000000000000000000000000000000000000000000000000", //Genesis Block always 0000
		HashMerkleRoot: "0000000000000000000000000000000000000000000000000000000000000000",
		TimeStamp:      TimeStamp(),
		Transaction_counter: 0,
		Transactions:        []*Transaction{},
		CreatorPubKey: creatorPubKey,
	}
	fmt.Println("New block :",block.BlockHash())
	b.Blocks = append(b.Blocks, block)
}


func (b *Blockchain) IsValidBlock(block *Block) bool {
	// Check if the block creator's public key is in the list of approved authorities
	for _, authority := range b.Authorities {
		if block.CreatorPubKey == authority {
			return true
		}
	}

	return false
}