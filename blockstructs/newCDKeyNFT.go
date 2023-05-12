package BlockStructs

import (
	"encoding/base64"
	"fmt"
)

func (b *Blockchain) NewNFTBlock(nfts []*CDKeyNFT, creatorPubKey string)  []*CDKeyNFT{

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
        Transaction_counter: len(b.Blocks),
        NFTTransactions:     []*NFTTransaction{},
        //NFTs:                nfts,
        CreatorPubKey:       creatorPubKey,
    }

    b.Blocks = append(b.Blocks, block)
    b.PendingNFTs = append(b.PendingNFTs, nfts...)
    //b.NewBlock(nil, b.Authorities[0]) //cahnge this to node address
    //b.Sync(b)
    return nfts
}




func ConvertToURLSafeBase64(input string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", fmt.Errorf("error decoding input string: %v", err)
	}

	urlSafeBase64 := base64.URLEncoding.EncodeToString(data)
	return urlSafeBase64, nil
}
