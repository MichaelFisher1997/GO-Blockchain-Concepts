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
			fmt.Printf("New Wallet Puublic Key: \n%s\n",base64.StdEncoding.EncodeToString(wallet.PublicKey))
			fmt.Println("--------------------")
			b.Wallets = append(b.Wallets, wallet)
}


//for API
func MakeWalletWithAmount(b *BlockStructs.Blockchain, initialAmount float64) (privateKeyStr string, publicKeyStr string, err error) {
	// Create a new wallet with the specified initial amount
	wallet := BlockStructs.NewWallet(initialAmount)

	// Convert the private and public keys to strings
	privateKeyStr = base64.StdEncoding.EncodeToString(wallet.PrivateKey)
	publicKeyStr = base64.StdEncoding.EncodeToString(wallet.PublicKey)

	// Add the wallet to the blockchain's wallet list
	b.Wallets = append(b.Wallets, wallet)


	return privateKeyStr, publicKeyStr, nil
}