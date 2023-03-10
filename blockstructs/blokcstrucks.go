package BlockStructs

import (
	"time"
)

type Transaction struct {
	// transaction fields
}

type Blockchain struct {
	Blocks []string//[]*Block test
	BlockCount int
	//Root  hash //Merkel root
	//root needs to loop through all the blocks and hash them all into a merkel root
}

type BlockHeader struct {
	Version        int    //Block version number, You upgrade the software and it specifies a new version, 4 bytles
	hashPrevBlock  string //256-bit hash of the previous block header, A new block comes in, 32 bytes
	hashMerkleRoot string //256-bit hash based on all of the transactions in the block, A transaction is accepted, 32 bytes
	TimeStamp 	   string //Current block timestamp as seconds since 1970-01-01T00:00 UTC, 4 bytes
}

type Block struct {
	BlockID				int // block number
	ThisBlockHash 			string //
	Magic_No            string     //value always 0xD9B4BEF9, 4 bytes
	Blocksize           int     //number of bytes following up to end of block, 4 bytes
	BlockHeader         BlockHeader     //consists of 6 items, 	80 bytes
	Transaction_counter int      // positive integer VI = VarInt, 1 - 9 bytes
	//Coinbase            string   //the first transaction in a block, 1 - 9 bytes, 32 characters
	Transactions        []Transaction // the (non empty) list of transaction, <Transaction counter>-many transactions
}

// fmt.Println(t.Format("20060102150405"))
func TimeStamp() string {
	return time.Now().Format(time.RFC850) //maybe change this to ow.UnixNano() // number of nanoseconds since January 1, 1970 UTC
}