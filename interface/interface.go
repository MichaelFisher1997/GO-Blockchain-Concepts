package Interface

import (
	"fmt"
	BlockStructs "go-blockchain/blockstructs"
	Read "go-blockchain/read"
	Utils "go-blockchain/utils"
)

type BlockInterface interface {
    Run()
}

func Run(b *BlockStructs.Blockchain) *BlockStructs.Blockchain{
	var i int
	for x := 0; x > -1; x++ {
		fmt.Print(
			"Welcome to Go-Blockchain\n",
			"select on of the following\n",
			"0 - quit: \n",
			"1 - new block: \n",
			"2 - Make transaction: \n ",
			"",
			"input number <- : ")
		fmt.Scan(&i)
	//--------------------------------------
		if i == 1 {
			b.NewBlock(nil)
			b.UpdateBalances()
			latestBlock := b.Blocks[len(b.Blocks)-1]
			fmt.Printf("Latest block: %v\n", latestBlock)
			err2 := Read.SaveToFile("ledger/myBlocks.json",b)
			Utils.Check(err2)
			// Save the current state of the blockchain to a binary file.
			if err := Read.SaveBlockchain(b, "blockchain.dat"); err != nil {
				fmt.Printf("Error saving blockchain: %v\n", err)
			}
			
		}
		if i == 0 {
			break
		}
	} 
	return b
}
