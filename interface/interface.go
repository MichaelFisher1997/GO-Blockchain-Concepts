package Interface

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/base64"
	"fmt"
	BlockStructs "go-blockchain/blockstructs"
	Commands "go-blockchain/commands"
	Read "go-blockchain/read"
	Utils "go-blockchain/utils"
	"math/big"
)

type BlockInterface interface {
    Run()
}

func Run(b *BlockStructs.Blockchain) *BlockStructs.Blockchain{
	//login or new wallet
	var x int
	fmt.Print("Welcome to Go-Blockchain\n",
		"1 - login: \n",
		"2 - new wallet: \n")
	fmt.Scan(&x)
	if x == 2 {
		Commands.MakeWallet(b)
	}
	//login
	// Ask for the private key
	var privateKeyStr string
	fmt.Print("Enter your private key: ")
	fmt.Scan(&privateKeyStr)

	// Decode the private key and derive the public key
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyStr)
	Utils.Check(err)
	curve := elliptic.P256()
	privateKey := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
		},
		D: new(big.Int).SetBytes(privateKeyBytes),
	}
	privateKey.PublicKey.X, privateKey.PublicKey.Y = curve.ScalarBaseMult(privateKey.D.Bytes())
	publicKeyBytes := elliptic.Marshal(curve, privateKey.PublicKey.X, privateKey.PublicKey.Y)
	publicKeyStr := base64.StdEncoding.EncodeToString(publicKeyBytes)

	var i int
	for x := 0; x > -1; x++ {
		fmt.Print(
			"select on of the following\n",
			"0 - quit: \n",
			"1 - new block: \n",
			"2 - Make transaction: \n",
			"3 - New Wallet: \n",
			"4 - Get Balance: \n",
			"5 - Make New NFT: \n",
			"6 - Get NFTS: \n",
			"7 - Make NFT Transaction Request: \n",
			"8 - confirm NFT Transaction: \n",
			"9 - Get NFT Transaction Requests: \n",
			"10 - Become Authority: \n",
			"input number <- : ")
		fmt.Scan(&i)
	//--------------------------------------
		if i == 1 {
			b.NewBlock(nil, publicKeyStr)
			//b.UpdateBalances()
			latestBlock := b.Blocks[len(b.Blocks)-1]
			fmt.Printf("Latest block: %v\n", latestBlock)
			Read.Sync(b)
			
		}
		if i == 2 {
			Commands.MakeTransaction(b)
			Read.Sync(b)
		}
		if i == 3 {
			Commands.MakeWallet(b)
			Read.Sync(b)
		}
		if i == 4 {
			var privateKeyStr string
			fmt.Print("Enter your private key: ")
			fmt.Scan(&privateKeyStr)
			Balance, err := Commands.GetBalanceByPrivateKey(b, privateKeyStr)
			Utils.Check(err)
			fmt.Println("____________________________________________________")
			fmt.Printf("Balance: %f\n", Balance)
			fmt.Println("____________________________________________________")
		}
		if 5 == i {
			var id uint64
			var cdKey string
			var tokenID uint64
			var privateKeyStr string
			fmt.Print("Enter the ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter the CD Key: ")
			fmt.Scanln(&cdKey)
			fmt.Print("Enter the Token ID: ")
			fmt.Scanln(&tokenID)
			fmt.Print("Enter your private key: ")
			fmt.Scanln(&privateKeyStr)
			// Decode the private key and derive the public key
			privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyStr)
			Utils.Check(err)
			curve := elliptic.P256()
			privateKey := &ecdsa.PrivateKey{
				PublicKey: ecdsa.PublicKey{
					Curve: curve,
				},
				D: new(big.Int).SetBytes(privateKeyBytes),
			}
			privateKey.PublicKey.X, privateKey.PublicKey.Y = curve.ScalarBaseMult(privateKey.D.Bytes())
			publicKeyBytes := elliptic.Marshal(curve, privateKey.PublicKey.X, privateKey.PublicKey.Y)
			nft, err := Commands.NewCDKeyNFT(id, cdKey, tokenID, publicKeyBytes)
			Utils.Check(err)
			b.NFTs = append(b.NFTs, nft)
			Read.Sync(b)
			//W8Q3M-GZ7LS-2D6TY-VK9PX-4H1FA
		}
		if 6 == i {
			var privateKeyStr string
			fmt.Print("Enter your private key: ")
			fmt.Scanln(&privateKeyStr)
			nft, err := Commands.ListOwnedNFTs(b, privateKeyStr)
			Utils.Check(err)
			for _, nft := range nft {
				fmt.Println("____________________________________________________")
				fmt.Printf("ID: %v\n", nft.ID)
				fmt.Printf("CD Key: %v\n", nft.CDKey)
				fmt.Printf("Token ID: %v\n", nft.TokenID)
				fmt.Printf("Owner: %v\n", nft.OwnerPubKey)
				fmt.Printf("Minted On: %v\n", nft.MintedOn)
				fmt.Println("____________________________________________________")
			}
		}
		if 7 == i {
			Commands.CreateNFTTransaction(b)
			Read.Sync(b)
		}
		if 8 == i {
			Commands.ConfirmNFTTransaction(b)
			Read.Sync(b)
		}
		if 10 == i {
			var privateKeyStr string
			fmt.Print("Enter your private key: ")
			fmt.Scanln(&privateKeyStr)
			BlockStructs.AddToAuthorities(b, privateKeyStr)
			Read.Sync(b)
		}
		if i == 0 {
			break
		}
	} 
	return b
}
