package Read

import (
	"encoding/json"
	blockstructs "go-blockchain/blockstructs"
	"os"
)

type MyStruct struct {
    MyArray []int
}

func SaveToFile(filename string, b *blockstructs.Blockchain) error {
    // Convert the struct to a JSON string
    jsonString, err := json.MarshalIndent(b,""," ")
    if err != nil {
        return err
    }

    // Write the JSON string to the file
    err = os.WriteFile(filename, jsonString, 0644)
    if err != nil {
        return err
    }

    return nil
}

func ReadFromFile(filename string) (*blockstructs.Blockchain, error) {
    // Read the contents of the file
    fileContents, err := os.ReadFile(filename)
    if err != nil {
        return &blockstructs.Blockchain{}, err
    }

    // Unmarshal the JSON string into a struct
    var myStruct MyStruct
    err = json.Unmarshal(fileContents, &myStruct)
    if err != nil {
        return &blockstructs.Blockchain{}, err
    }

    return &blockstructs.Blockchain{}, nil
}

