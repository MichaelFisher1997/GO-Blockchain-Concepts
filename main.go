package main

import (
	"fmt"
	BlockStructs "go-blockchain/blockstructs"
	"os"
)
func writeToFile(filename string, data string) error {
	// Open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the data to the file
	_, err = fmt.Fprintln(file, data)
	if err != nil {
		return err
	}
	return nil
}

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
	Data := BlockStructs.BlockJSON(blockchain.Blocks[0])
	writeToFile("blockchain.txt", Data)

}

