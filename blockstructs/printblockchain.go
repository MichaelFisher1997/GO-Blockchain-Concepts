package BlockStructs

import (
	"fmt"
	//Utils "go-blockchain/utils"
	//"os"
	//"strconv"
)

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
/* testing!!!
func BlockPrint(b *Blockchain) {
	f, err := os.Create("./ledger/data.txt")
	Utils.Check(err)
	defer f.Close()
	blocks := b.Blocks
	for _, blocks := range b.Blocks {
		spacer := "_____________________________"
		blockID := blocks.BlockID
		headerVersion := blocks.BlockHeader.Version
		hashMerkleRoot := blocks.BlockHeader.hashMerkleRoot
		hashPrevBlock := blocks.BlockHeader.hashPrevBlock
		MagicNo := blocks.Magic_No
		TransactionCount := blocks.Transaction_counter
		TimeStamp := blocks.BlockHeader.TimeStamp

		_, err2 := f.WriteString(string(spacer)+"\n")
		Utils.Check(err2)
		_, err3 := f.WriteString("BlockID: " + strconv.Itoa(blockID)+"\n")
		Utils.Check(err3)
		_, err4 := f.WriteString("Version: " + strconv.Itoa(headerVersion)+"\n")
		Utils.Check(err4)
		_, err5 := f.WriteString("HashMerkleRoot: " + string(hashMerkleRoot)+"\n")
		Utils.Check(err5)
		_, err6 := f.WriteString("PrevBlockHash: " + hashPrevBlock+"\n")
		Utils.Check(err6)
		_, err7 := f.WriteString("MagicNo: " + MagicNo+"\n")
		Utils.Check(err7)
		_, err8 := f.WriteString("Transaction Count: " + strconv.Itoa(TransactionCount)+"\n")
		Utils.Check(err8)
		_, err9 := f.WriteString("Time Stamp: " + TimeStamp+"\n")
		Utils.Check(err9)
	}
	fmt.Printf("blocks: %v\n", blocks)
}*/