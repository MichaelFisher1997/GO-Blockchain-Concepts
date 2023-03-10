package BlockStructs

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	Utils "go-blockchain/utils"
)

func (b *Blockchain) NewBlock() {
	b.BlockCount = b.BlockCount+1
	PrevBlockHash :=  ""

	block := &Block{
		Magic_No: "0xF9S834SK",
		BlockID: b.BlockCount,
		Blocksize: 80,
		BlockHeader: BlockHeader{
			Version:        1,
			hashPrevBlock:  PrevBlockHash,
			hashMerkleRoot: b.MerkelRoot(),
			TimeStamp:      TimeStamp(),
		},
		Transaction_counter: 1,
		Transactions:        []Transaction{},
	}
	blockHash := block.BlockHash()
	block.ThisBlockHash = blockHash
	new, err := json.MarshalIndent(block, "", " ")
	Utils.Check(err)
	b.Blocks = append(b.Blocks, string(new))
}

func (b *Blockchain) NewGenesisBlock() string{
	block := &Block{
		Magic_No: "0xD9B4BEF9",
		BlockID: 0,
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
	block.ThisBlockHash = blockHash
	new, err := json.MarshalIndent(block, "", " ")
	Utils.Check(err)
	//b.Blocks = append(b.Blocks, string(new))
	return string(new)
}



func (b *Block) BlockHash() string {
	// trunk-ignore(golangci-lint/govet)
	header := fmt.Sprintf("%d%s%s%d" ,b.BlockHeader.Version, b.BlockHeader.hashPrevBlock, b.BlockHeader.hashMerkleRoot, b.BlockHeader.TimeStamp)
	h := sha256.Sum256([]byte(header))
	return hex.EncodeToString(h[:])
}

func NewBlockchain() *Blockchain {
	blockchain := &Blockchain{}
	genesisBlock := blockchain.NewGenesisBlock()
	blockchain.Blocks = append(blockchain.Blocks, string(genesisBlock))
	return blockchain
}

