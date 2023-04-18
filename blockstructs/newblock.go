package BlockStructs

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func (b *Blockchain) NewBlock(transactions []*Transaction) {
	b.BlockCount = b.BlockCount+1
	//PrevBlockHash :=  ""

	block := &Block{
		Magic_No: "0xF9S834SK",
		BlockID: b.BlockCount,
		Blocksize: 80,
		Version:		1,
		HashPrevBlock:  b.PrevBlockHash(),
		HashMerkleRoot: b.MerkelRoot(),
		TimeStamp:      TimeStamp(),
		Transaction_counter: len(transactions),
        Transactions:        transactions,
	}
	fmt.Println("New block :",block.BlockHash())
	b.Blocks = append(b.Blocks, block)
	// Clear the pending transactions after adding them to the new block
	b.PendingTransactions = []*Transaction{}
}

func (b *Blockchain) NewNFTBlock(nftTransactions []*NFTTransaction) {
	b.BlockCount = b.BlockCount + 1

	block := &Block{
		Magic_No:          "0xF9S834SK",
		BlockID:           b.BlockCount,
		Blocksize:         80,
		Version:           1,
		HashPrevBlock:     b.PrevBlockHash(),
		HashMerkleRoot:    b.MerkelRoot(),
		TimeStamp:         TimeStamp(),
		Transaction_counter: len(nftTransactions),
		NFTTransactions:     nftTransactions,
	}
	// Update wallet balances and NFT ownership
	for _, nftTransaction := range nftTransactions {
		if nftTransaction.Confirmed {
			// Deduct the amount from the seller's wallet
			senderPublicKeyBytes, err := base64.StdEncoding.DecodeString(nftTransaction.SenderPubKey)
			if err != nil {
				fmt.Println("Error decoding seller's public key:", err)
				return
			}
			sellerWallet := b.findWalletByPublicKey(senderPublicKeyBytes)
			if sellerWallet != nil {
				sellerWallet.Balance -= nftTransaction.Amount
			}

			// Decode the buyer's public key from base64
			receiverPublicKeyBytes, err := base64.StdEncoding.DecodeString(nftTransaction.ReceiverPubKey)
			if err != nil {
				fmt.Println("Error decoding buyer's public key:", err)
				return
			}

			// Add the amount to the buyer's wallet
			buyerWallet := b.findWalletByPublicKey(receiverPublicKeyBytes)
			if buyerWallet != nil {
				buyerWallet.Balance += nftTransaction.Amount
			}

			// Update the NFT's ownership
			nft := b.findNFTByID(nftTransaction.NFTID)
			if nft != nil {
				nft.OwnerPubKey = nftTransaction.ReceiverPubKey
			}
		}
	}
	fmt.Println("New NFT block :", block.BlockHash())
	b.Blocks = append(b.Blocks, block)
	// Clear the pending NFT transactions after adding them to the new block
	b.PendingNFTTransactions = []*NFTTransaction{}
}







func (b *Block) BlockHash() string {
	//convert to string
	header := fmt.Sprintf("%v" ,b)
	h := sha256.Sum256([]byte(header))
	return hex.EncodeToString(h[:])
}

func NewBlockchain() *Blockchain {
    blockchain := &Blockchain{
        Wallets: []*Wallet{},
    }

    // Add the genesis block to the blockchain
    blockchain.NewGenesisBlock()

    return blockchain
}

/*func (b *Blockchain) AddTransaction(transaction *Transaction) {
	// Add the transaction to the list of pending transactions
	b.PendingTransactions = append(b.PendingTransactions, transaction)

	// If there are enough transactions in the pending transactions list,
	// create a new block.
	// This value can be adjusted based on the desired number of transactions per block.
	if len(b.PendingTransactions) >= 1 {
		b.NewBlock()
	}
}*/

func (b *Blockchain) NewGenesisBlock() {
	block := &Block{
		Magic_No: "0xD9B4BEF9",
		BlockID: 0,
		Blocksize: 80,
		Version:        1, //You upgrade the software and it specifies a new version
		HashPrevBlock:  "0000000000000000000000000000000000000000000000000000000000000000", //Genesis Block always 0000
		HashMerkleRoot: "0000000000000000000000000000000000000000000000000000000000000000",
		TimeStamp:      TimeStamp(),
		Transaction_counter: 0,
		Transactions:        []*Transaction{},
	}
	fmt.Println("New block :",block.BlockHash())
	b.Blocks = append(b.Blocks, block)
}