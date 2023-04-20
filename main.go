package main

import (
	"fmt"
	BlockStructs "go-blockchain/blockstructs"

	//Interface "go-blockchain/interface"
	api "go-blockchain/api"
	Read "go-blockchain/read"
	Utils "go-blockchain/utils"
)

func main() {
	fmt.Println("------start------")
	blockchain, err := Read.LoadBlockchain("blockchain.dat")
	if err != nil {
		// If there is an error (e.g., the file doesn't exist), create a new blockchain with the sender and receiver wallets.
        //blockchain = BlockStructs.NewBlockchain()
		blockchain = BlockStructs.Start()
	}
	//apitest
	//blockchain.NewBlock(nil,blockchain.Authorities[0])
	//Interface.Run(blockchain)
	Read.Sync(blockchain)
	api.ApiRun(blockchain)

	//----------------SAVE-------------
	//time.Sleep(1 * time.Second)
	err2 := Read.SaveToFile("ledger/myBlocks.json",blockchain)
	Utils.Check(err2)
	// Save the current state of the blockchain to a binary file.
	if err := Read.SaveBlockchain(blockchain, "blockchain.dat"); err != nil {
		fmt.Printf("Error saving blockchain: %v\n", err)
	}
}