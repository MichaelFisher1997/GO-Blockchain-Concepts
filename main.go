package main

import (
	"fmt"
	BlockStructs "go-blockchain/blockstructs"
)


func main() {
	fmt.Println("------start------")

	blockchain := BlockStructs.NewBlockchain()
	blockchain.Blocks = append(blockchain.Blocks, blockchain.NewBlock())
	blockchain.Blocks = append(blockchain.Blocks, blockchain.NewBlock())
	blockchain.Blocks = append(blockchain.Blocks, blockchain.NewBlock())
	fmt.Println("Blockchain: ", blockchain.Blocks)
	fmt.Println("Merkel Root: ", blockchain.MerkelRoot())
	fmt.Println("-------------------------------")
	BlockStructs.PrintBlock(blockchain.Blocks[0])
	BlockStructs.PrintBlock(blockchain.Blocks[1])
	BlockStructs.PrintBlock(blockchain.Blocks[2])
	BlockStructs.PrintBlock(blockchain.Blocks[3])

}

