package BlockStructs

import (
	"crypto/sha256"
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
}

func (b *Block) BlockHash() string {
	//convert to string
	header := fmt.Sprintf("%v" ,b)
	h := sha256.Sum256([]byte(header))
	return hex.EncodeToString(h[:])
}

func NewBlockchain() *Blockchain {
    blockchain := &Blockchain{
        Wallets: []*Wallet{},
		Authorities: []string{},
    }


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