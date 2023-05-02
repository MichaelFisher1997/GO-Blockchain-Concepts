package BlockStructs

import (
	"encoding/base64"
	"fmt"
	//BlockStructs "go-blockchain/blockstructs"
	//Blockchain "blockchain/blockchain"
)

func NewNFTBlock(b *Blockchain, nfts []*CDKeyNFT, creatorPubKey string) error {
    tempBlock := &Block{
        CreatorPubKey: creatorPubKey,
    }
    if !b.IsValidBlock(tempBlock) {
        return fmt.Errorf("error: block creator is not an authorized authority")
    }

    b.BlockCount = b.BlockCount + 1
    fmt.Print(creatorPubKey)

    block := &Block{
        Magic_No:          "0xF9S834SK",
        BlockID:           b.BlockCount,
        Blocksize:         80,
        Version:           1,
        HashPrevBlock:     b.PrevBlockHash(),
        HashMerkleRoot:    b.MerkelRoot(),
        TimeStamp:         TimeStamp(),
        Transaction_counter: 0,
        NFTTransactions:     []*NFTTransaction{},
        //NFTs:                nfts,
        CreatorPubKey:       creatorPubKey,
    }

    b.Blocks = append(b.Blocks, block)
    b.PendingNFTs = append(b.PendingNFTs, nfts...)
    b.NewBlock(nil, creatorPubKey)
    return nil
}




func ConvertToURLSafeBase64(input string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", fmt.Errorf("error decoding input string: %v", err)
	}

	urlSafeBase64 := base64.URLEncoding.EncodeToString(data)
	return urlSafeBase64, nil
}

/*func AddNFTBlock(b *BlockStructs.Blockchain, nftTransactions []*BlockStructs.NFTTransaction, creatorPubKey string) error {
	err := NewNFTBlock(b, creatorPubKey)
	if err != nil {
		return err
	}

	return nil
}


/*
input number <- : 5
Enter the ID: 123
Enter the CD Key: AB12-CD34-EF56-GH78
Enter the Token ID: 1001
Enter your private key: WgFXmDQZ8aHtGmRjI9X9trLmPp8WxjKq1cbrDtI7Npk=
*/
