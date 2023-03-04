package BlockStructs

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func (b *Blockchain) NewBlock() *Block {
	block := &Block{
		Magic_No: "0xF9S834SK",
		Blocksize: 80,
		BlockHeader: BlockHeader{
			Version:        1,
			hashPrevBlock:  b.PrevBlockHash(),
			hashMerkleRoot: b.MerkelRoot(),
			TimeStamp:      TimeStamp(),
		},
		Transaction_counter: 1,
		Transactions:        []Transaction{},
	}
	blockHash := block.BlockHash()
	fmt.Println("New Block Hash: ", blockHash)
	return block
}

func NewGenesisBlock() *Block {
	block := &Block{
		Magic_No: "0xD9B4BEF9",
		Blocksize: 80,
		BlockHeader: BlockHeader{
			Version:        1, //You upgrade the software and it specifies a new version
			hashPrevBlock:  "0000000000000000000000000000000000000000000000000000000000000000", //Genesis Block always 0000
			hashMerkleRoot: "0000000000000000000000000000000000000000000000000000000000000000",
			TimeStamp:      TimeStamp(),
		},
		Transaction_counter: 1,
		Transactions:        []Transaction{},
	}
	blockHash := block.BlockHash()
	fmt.Println("Genesis Block Hash: ", blockHash)
	return block
}



func (b *Block) BlockHash() string {
	// trunk-ignore(golangci-lint/govet)
	header := fmt.Sprintf("%d%s%s%d" ,b.BlockHeader.Version, b.BlockHeader.hashPrevBlock, b.BlockHeader.hashMerkleRoot, b.BlockHeader.TimeStamp)
	h := sha256.Sum256([]byte(header))
	return hex.EncodeToString(h[:])
}

func NewBlockchain() *Blockchain {
	blockchain := &Blockchain{}
	genesisBlock := NewGenesisBlock()
	blockchain.Blocks = []*Block{genesisBlock}
	return blockchain
}

