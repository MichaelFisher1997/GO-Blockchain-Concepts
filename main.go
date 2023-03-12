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

