package api

import (
	"encoding/base64"
	"fmt"
	BlockStructs "go-blockchain/blockstructs"
	Commands "go-blockchain/commands"

	//Interface "go-blockchain/interface"
	Read "go-blockchain/read"
	Utils "go-blockchain/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ApiRun(b *BlockStructs.Blockchain) {
	// Set up Gin router
	router := gin.Default()

	 // Add CORS middleware
	 router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000", "http://0.0.0.0:3000"},
        AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

	//testing routers
	router.GET("/clear_pending", func(c *gin.Context) {
		// Clear the pending transactions
		b.PendingTransactions = []*BlockStructs.Transaction{}
		b.PendingNFTTransactions = []*BlockStructs.NFTTransaction{}
		b.PendingNFTs = []*BlockStructs.CDKeyNFT{}
		c.JSON(http.StatusOK, gin.H{
			"message": "Pending transactions cleared",
		})
		Read.Sync(b)
	})

	router.GET("/mine", func(c *gin.Context) {
		// Create a new block with the pending transactions
		b.Mine()
		c.JSON(http.StatusOK, gin.H{
			"message": "New block successfully mined",
		})
		Read.Sync(b)
	})
	//end testing routers


	router.GET("/create_wallet", func(c *gin.Context) {
		// Get the initial amount from the request query parameter
		initialAmountStr := c.Query("initial_amount")
		initialAmount, err := strconv.ParseFloat(initialAmountStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid initial_amount"})
			return
		}

		// Call the MakeWalletWithAmount function and return the result
		privateKeyStr, publicKeyStr, err := Commands.MakeWalletWithAmount(b, initialAmount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"private_key": privateKeyStr,
			"public_key":  publicKeyStr,
		})
		Read.Sync(b)
	})
	//http://localhost:8080/create_wallet?initial_amount=75

	router.GET("/make_transaction", func(c *gin.Context) {
		// Get the private key, recipient's public key, and amount from the URL query parameters
		privateKeyStr := c.Query("private_key")
		recipientPublicKeyStr := c.Query("recipient_public_key")
		amountStr := c.Query("amount")

		// Convert the amount string to float64
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
			return
		}

		// Call the MakeTransaction function with the provided data
		err = Commands.MakeTransactionWithDetails(b, privateKeyStr, recipientPublicKeyStr, amount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Transaction successfully created and added to the pending transactions",
		})
		Read.Sync(b)
	})
	//http://localhost:8080/make_transaction?private_key=P4kXX2bGQdZXkAf7WuPUehsqqf4kktsJXWeXGSdEdr4%3D&recipient_public_key=BOcbcq8c0XYVbCRuQSJCfwEsIW%2Fxy9sjhXf8g67w%2FBDeFxQrOXPuzjCtJo4LdtvfgmghOKnugKQ6ki7bb%2BOam%2BU%3D&amount=10

	// Add other API endpoints for your blockchain commands
	// ...
	router.GET("/login", func(c *gin.Context) {
		// Get the URL-encoded private key from the URL query parameters
		privateKeyStr := c.Query("private_key")
		fmt.Print("PK: ", privateKeyStr, "\n")
		// Call the GetBalanceByPrivateKey function with the decoded private key
		balance, publicKeyStr, err := Commands.GetBalanceByPrivateKey(b, privateKeyStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Get the list of owned NFTs
		ownedNFTs:= Commands.ListOwnedNFTs(b, privateKeyStr)
		Utils.Check(err)

		// Find the corresponding wallet in the blockchain
		var wallet *BlockStructs.Wallet
		for _, w := range b.Wallets {
			if publicKeyStr == base64.StdEncoding.EncodeToString(w.PublicKey) {
				wallet = w
				break
			}
		}
		if wallet == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Wallet not found"})
			return
		}

		// Return the public key, balance, and NFTs
		c.JSON(http.StatusOK, gin.H{
			"public_key": publicKeyStr,
			"balance":    balance,
			"nfts":       ownedNFTs,
		})
	})

	//Create NFT
	router.GET("/create_nft", func(c *gin.Context) {
		idStr := c.Query("id")
		cdKey := c.Query("cd_key")
		tokenIDStr := c.Query("token_id")
		creatorPrivateKey := c.Query("private_key")
		//
		//get public key
		//fmt.Print("PK: ", creatorPrivateKey, "\n")
		creatorPubKey, err := Utils.DecodePrivateKey(creatorPrivateKey)
		Utils.Check(err)
		//fmt.Fprintln(c.Writer, "creatorPubKey: ", creatorPubKey)

		// Convert id and tokenID to uint64
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}

		nft := &BlockStructs.CDKeyNFT{
			ID:          id,
			CDKey:       cdKey,
			TokenID:     tokenIDStr,
			Minted:      false,
			MintedBy:   creatorPubKey,
			MintedOn:   BlockStructs.TimeStamp(),
			OwnerPubKey: creatorPubKey,
		}
		// Call NewNFTBlock function
		New_nft := b.NewNFTBlock([]*BlockStructs.CDKeyNFT{nft}, creatorPubKey)
		//Utils.Check(New_nft)

		c.JSON(http.StatusOK, gin.H{
			"message": "NFT added to pending NFTs",
			"nft":     New_nft,
		})
		Read.Sync(b)
	})
	//----------test freom down-------
	router.GET("/make_nft_transaction", func(c *gin.Context) {
		// Get the private key, recipient's public key, and amount from the URL query parameters
		nftIDStr := c.Query("nft_id")
		amountStr := c.Query("amount")
		privateKeyStr := c.Query("private_key")

		// Convert the amount string to float64
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
			return
		}

		// Convert the nftID string to uint64
		nftID, err := strconv.ParseUint(nftIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid nft_id"})
			return
		}

		// Call the MakeTransaction function with the provided data
		err = b.PendingNFTTransaction(privateKeyStr, nftID, amount)
		Utils.Check(err)

		c.JSON(http.StatusOK, gin.H{
			"message": "NFT transaction successfully created and added to the pending transactions",
		})
		Read.Sync(b)
	})

	//List NFT adds
	router.GET("/list_nft_adds", func(c *gin.Context) {
		adds := b.AllPendingNFTTransactions()
		c.JSON(http.StatusOK, gin.H{
			"adds": adds,
		})
	})

	//buy NFT
	router.GET("/buy_nft", func(c *gin.Context) {
		// Get the private key, recipient's public key, and amount from the URL query parameters
		nftIDStr := c.Query("nft_id")
		amountStr := c.Query("amount")
		privateKeyStr := c.Query("private_key")

		// Convert the amount string to float64
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
			return
		}

		// Convert the nftID string to uint64
		nftID, err := strconv.ParseUint(nftIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid nft_id"})
			return
		}

		// Call the MakeTransaction function with the provided data
		err = b.BuyNFT(nftID, privateKeyStr, amount)
		Utils.Check(err)

		c.JSON(http.StatusOK, gin.H{
			"message": "NFT transaction successfully created and added to the pending transactions",
		})
		Read.Sync(b)
	})

	
	// Start the API server
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}

}