package Read

import (
	"fmt"
	"os"
	"path/filepath"
)

func Findledger() {
    filename := "data.txt"
	dir := "ledger"
	foundPath, err := findFile(filename, dir)
	if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Found file %s at path %s\n", filename, foundPath)
}

func findFile(filename string, dir string) (string, error) {
    // Get the absolute path of the directory
    absDir, err := filepath.Abs(dir)
    if err != nil {
        return "", err
    }

    // Walk the directory tree starting from the absolute path
    var foundPath string
    err = filepath.Walk(absDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && info.Name() == filename {
            // Found the file
            foundPath = path
            return filepath.SkipDir
        }
        return nil
    })
    if err != nil {
        return "", err
    }
    if foundPath == "" {
        return "", fmt.Errorf("file not found")
    }
    return foundPath, nil
}