package BlockStructs

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func (b *Blockchain) MerkelRoot() string {
	hashes := []string{}
	for _, block := range b.Blocks {
		blockHash,err := json.Marshal(block)
		if err != nil {
			panic(err)
		}
		hashed := sha256.Sum256(blockHash)
		hashString := hex.EncodeToString(hashed[:])
		hashes = append(hashes, hashString)
	}

	// Duplicate the last hash if the array has an odd number of elements
    if len(hashes)%2 != 0 {
        hashes = append(hashes, hashes[len(hashes)-1])
    }

    // Calculate the Merkle root by hashing pairs of hashes
    for len(hashes) > 1 {
		var newHashes []string
		for i := 0; i < len(hashes); i += 2 {
			hash1, _ := hex.DecodeString(hashes[i])
			var hash2 []byte
			if i+1 < len(hashes) {
				hash2, _ = hex.DecodeString(hashes[i+1])
			} else {
				hash2 = hash1
			}
			concatenated := append(hash1, hash2...)
			hash := sha256.Sum256(concatenated)
			newHashes = append(newHashes, hex.EncodeToString(hash[:]))
		}
		hashes = newHashes
	}
	return string(hashes[0])
}

func (b *Blockchain) PrevBlockHash() string {
	//last := []string{}
	if  len(b.Blocks) <= 0 {
		//convert to string
		last := fmt.Sprintf("%v" , (b.Blocks[1]))
		hash := sha256.Sum256([]byte(last))
	return hex.EncodeToString(hash[:])
	} else {
		//convert to string
		last := fmt.Sprintf("%v" , b.Blocks[len(b.Blocks)-1])
		hash := sha256.Sum256([]byte(last))
	return hex.EncodeToString(hash[:])
	}
}