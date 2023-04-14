package Read

import (
	"fmt"
	BlockStructs "go-blockchain/blockstructs"
	Utils "go-blockchain/utils"
)

func Sync(b *BlockStructs.Blockchain) {
	err := SaveToFile("ledger/myBlocks.json", b)
	Utils.Check(err)
	// Save the current state of the blockchain to a binary file.
	if err := SaveBlockchain(b, "blockchain.dat"); err != nil {
		fmt.Printf("Error saving blockchain: %v\n", err)
	}
}