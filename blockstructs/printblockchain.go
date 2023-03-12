package BlockStructs

import (
	"fmt"
)

//redundent!!
func PrintBlock(block *Block) {
	fmt.Println("-------------------------------")
	fmt.Println("Block:")
	fmt.Printf("  Magic Number: %s\n", block.Magic_No)
	fmt.Printf("  Block Size: %d bytes\n", block.Blocksize)

	fmt.Println("  Block Header:")
	fmt.Printf("    Version: %d\n", block.Version)
	fmt.Printf("    Previous Block Hash: %s\n", block.HashPrevBlock)
	fmt.Printf("    Merkle Root Hash: %s\n", block.HashMerkleRoot)
	fmt.Printf("    Timestamp: %s\n", block.TimeStamp)

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