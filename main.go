package main

import (
	"fmt"
	BlockStructs "go-blockchain/blockstructs"
	Read "go-blockchain/read"
	"time"
	//"encoding/json"
)

func main() {
	fmt.Println("------start------")

	blockchain := BlockStructs.NewBlockchain()
	blockchain.NewBlock()
	blockchain.NewBlock()
	fmt.Print(blockchain.Blocks[0].BlockHash(), "\n")
	//blockchain.Blocks = append(blockchain.Blocks, json.MarshalIndent(blockchain.NewBlock(), "", "  "))
	//blockchain.Blocks = append(blockchain.Blocks, blockchain.NewBlock())
	//blockchain.Blocks = append(blockchain.Blocks, blockchain.NewBlock())
	//fmt.Println("Blockchain: ", blockchain.Blocks)
	//fmt.Println("Merkel Root: ", blockchain.MerkelRoot())
	//fmt.Println("-------------------------------")
	//BlockStructs.PrintBlock(blockchain.Blocks[0])
	//BlockStructs.PrintBlock(blockchain.Blocks[1])
	//BlockStructs.PrintBlock(blockchain.Blocks[2])
	//BlockStructs.PrintBlock(blockchain.Blocks[3])
	//Data := BlockStructs.BlockJSON(blockchain.Blocks[0])
	//fmt.Println(Data)
	//BlockStructs.BlockPrint(blockchain)
	time.Sleep(1 * time.Second)
	Read.Findledger()
	err := Read.SaveToFile("ledger/myBlocks.json", blockchain)
	if err != nil {
        fmt.Println(err)
        return
    }
	mySavedBlockchain, err := Read.ReadFromFile("ledger/myBlocks.json")
	if err != nil {
        fmt.Println(err)
        return
    }
	fmt.Printf("Array in file: %v\n", mySavedBlockchain)
}

