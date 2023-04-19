package BlockStructs

import (
	"time"
)

type Wallet struct {
	PublicKey  []byte
	PrivateKey []byte
	Balance    float64
}

type Transaction struct {
	SenderPublicKey    []byte
	ReceiverPublicKey  []byte
	Amount             float64
	Signature          []byte
}

type CDKeyNFT struct {
    ID          uint64 `json:"id"`
    CDKey       string `json:"cd_key"`
    TokenID     uint64 `json:"token_id"`
    Minted      bool   `json:"minted"`
    MintedBy    string `json:"minted_by"`
    MintedOn    string `json:"minted_on"`
	OwnerPubKey string `json:"owner_pub_key"`
}

type NFTTransaction struct {
	ID             uint64 `json:"id"`
	NFTID          uint64 `json:"nft_id"`
	SenderPubKey   string `json:"sender_pub_key"`
	ReceiverPubKey string `json:"receiver_pub_key"`
	Amount         float64 `json:"amount"`
	Signature      string `json:"signature"`
	Confirmed      bool   `json:"confirmed"`
}

type Blockchain struct {
	Blocks []*Block
	BlockCount int
	Wallets []*Wallet
	PendingTransactions []*Transaction
	PendingNFTTransactions []*NFTTransaction `json:"pending_nft_transactions"`
	NFTs []*CDKeyNFT
	Authorities          []string
	//Root  hash //Merkel root
	//root needs to loop through all the blocks and hash them all into a merkel root
}

type Block struct {
	//Header
	Version        		int    //Block version number, You upgrade the software and it specifies a new version, 4 bytles
	HashPrevBlock  		string //256-bit hash of the previous block header, A new block comes in, 32 bytes
	HashMerkleRoot 		string //256-bit hash based on all of the transactions in the block, A transaction is accepted, 32 bytes
	TimeStamp 	   		string //Current block timestamp as seconds since 1970-01-01T00:00 UTC, 4 bytes
	BlockID				int // block number
	Magic_No            string     //value always 0xD9B4BEF9, 4 bytes
	Blocksize           int     //number of bytes following up to end of block, 4 bytes
	Transaction_counter int      // positive integer VI = VarInt, 1 - 9 bytes
	//Coinbase            string   //the first transaction in a block, 1 - 9 bytes, 32 characters
	Transactions 		[]*Transaction // the (non empty) list of transaction, <Transaction counter>-many transactions
	NFTTransactions []*NFTTransaction `json:"nft_transactions"`
	CreatorPubKey string

}

// fmt.Println(t.Format("20060102150405"))
func TimeStamp() string {
	return time.Now().Format(time.RFC850) //maybe change this to ow.UnixNano() // number of nanoseconds since January 1, 1970 UTC
}