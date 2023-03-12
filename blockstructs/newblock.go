package BlockStructs

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func (b *Blockchain) NewBlock() {
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
		Transaction_counter: 0,
		Transactions:        []Transaction{},
	}
	fmt.Println("New block :",block.BlockHash())
	//block.ThisBlockHash = blockHash
	//new, err := json.MarshalIndent(block, "", " ")
	//Utils.Check(err)
	b.Blocks = append(b.Blocks, block)
}

func (b *Blockchain) NewGenesisBlock() {
	block := &Block{
		Magic_No: "0xD9B4BEF9",
		BlockID: 0,
		Blocksize: 80,
		Version:        1, //You upgrade the software and it specifies a new version
		HashPrevBlock:  "0000000000000000000000000000000000000000000000000000000000000000", //Genesis Block always 0000
		HashMerkleRoot: "0000000000000000000000000000000000000000000000000000000000000000",
		TimeStamp:      TimeStamp(),
		Transaction_counter: 0,
		Transactions:        []Transaction{},
	}
	fmt.Println("New block :",block.BlockHash())
	//block.ThisBlockHash = blockHash
	//new, err := json.MarshalIndent(block, "", " ")
	//Utils.Check(err)
	b.Blocks = append(b.Blocks, block)
}



func (b *Block) BlockHash() string {
	//convert to string
	header := fmt.Sprintf("%v" ,b)
	h := sha256.Sum256([]byte(header))
	return hex.EncodeToString(h[:])
}

func NewBlockchain() *Blockchain {
	blockchain := &Blockchain{}
	//genesisBlock := blockchain.NewGenesisBlock()
	blockchain.NewGenesisBlock()
	//blockchain.Blocks = append(blockchain.Blocks, genesisBlock)
	return blockchain
}

