package BlockStructs

import (
	"encoding/json"
	"fmt"
)

func PrintBlock(block *Block) {
	fmt.Println("-------------------------------")
	fmt.Println("Block:")
	fmt.Printf("  Magic Number: %s\n", block.Magic_No)
	fmt.Printf("  Block Size: %d bytes\n", block.Blocksize)

	header := block.BlockHeader
	fmt.Println("  Block Header:")
	fmt.Printf("    Version: %d\n", header.Version)
	fmt.Printf("    Previous Block Hash: %s\n", header.hashPrevBlock)
	fmt.Printf("    Merkle Root Hash: %s\n", header.hashMerkleRoot)
	fmt.Printf("    Timestamp: %s\n", header.TimeStamp)

	fmt.Printf("  Transaction Counter: %d\n", block.Transaction_counter)

	if len(block.Transactions) == 0 {
		fmt.Println("  No transactions in block")
	} else {
		fmt.Println("  Transactions:")
		/*  Add later!!!
		for i, tx := range block.Transactions {
			tx = nullTx
			fmt.Printf(tx)
			fmt.Printf("    Transaction %d:\n", i+1)
			// Print out transaction fields as desired
		}
		*/
	}
}

func BlockJSON(b *Block) string {
	// Convert the block to a JSON string
	blockJSON, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return ""
	}

	return string(blockJSON)
}
