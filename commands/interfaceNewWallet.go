package Commands

import (
	"encoding/base64"
	"fmt"
	BlockStructs "go-blockchain/blockstructs"
)

func MakeWallet(b *BlockStructs.Blockchain) {
	var number float64
			fmt.Print("Enter Wallet amount: ")
			_, err := fmt.Scanf("%f", &number)
			if err != nil {
				fmt.Println("Error reading input:", err)
			}
			wallet := BlockStructs.NewWallet(number)
			fmt.Println("--------------------")
			fmt.Printf("New Wallet Private Key: \n%s\n",base64.StdEncoding.EncodeToString(wallet.PrivateKey))
			fmt.Printf("New Wallet Private Key: \n%s\n",base64.StdEncoding.EncodeToString(wallet.PublicKey))
			fmt.Println("--------------------")
			b.Wallets = append(b.Wallets, wallet)
}