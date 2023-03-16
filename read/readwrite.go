package Read

import (
	"encoding/json"
	BlockStructs "go-blockchain/blockstructs"
	blockstructs "go-blockchain/blockstructs"
	"os"

	//"os"
	"bytes"
	"encoding/gob"
)

func SaveToFile(filename string, b *blockstructs.Blockchain) error {
    // Convert the struct to a JSON string
    jsonString, err := json.MarshalIndent(b,"", " ")
    if err != nil {
        return err
    }

    // Write the JSON string to the file
    err = os.WriteFile(filename, jsonString, 0644)
    if err != nil {
        panic(err)
    }

    return nil
}

// SaveBlockchain saves the current state of the blockchain to a binary file.
func SaveBlockchain(blockchain *BlockStructs.Blockchain, filename string) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(blockchain); err != nil {
		return err
	}

	if err := os.WriteFile(filename, buf.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}

// LoadBlockchain reads and loads the saved blockchain state from a binary file.
func LoadBlockchain(filename string) (*BlockStructs.Blockchain, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var blockchain BlockStructs.Blockchain
	buf := bytes.NewReader(data)
	dec := gob.NewDecoder(buf)

	if err := dec.Decode(&blockchain); err != nil {
		return nil, err
	}

	return &blockchain, nil
}