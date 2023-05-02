package api

import (
	"encoding/base64"
	BlockStructs "go-blockchain/blockstructs"
	Commands "go-blockchain/commands"
	Read "go-blockchain/read"
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
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

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
		// Get the private key from the URL query parameters
		privateKeyStr := c.Query("private_key")

		// Call the GetBalanceByPrivateKey function with the provided private key
		balance, publicKeyStr, err := Commands.GetBalanceByPrivateKey(b, privateKeyStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Get the list of owned NFTs
		ownedNFTs, err := Commands.ListOwnedNFTs(b, privateKeyStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
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

	router.GET("/create_nft", func(c *gin.Context) {
		idStr := c.Query("id")
		cdKey := c.Query("cd_key")
		tokenIDStr := c.Query("token_id")
		creatorPubKey := c.Query("creator_public_key")

		// Convert id and tokenID to uint64
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}

		tokenID, err := strconv.ParseUint(tokenIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token ID"})
			return
		}

		nft := &BlockStructs.CDKeyNFT{
			ID:          id,
			CDKey:       cdKey,
			TokenID:     tokenID,
			Minted:      false,
			OwnerPubKey: creatorPubKey,
		}


		c.JSON(http.StatusOK, gin.H{
			"message": "NFT added to pending NFTs",
			"nft":     nft,
		})
	})
	



	// Start the API server
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}

}
