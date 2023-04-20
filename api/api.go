package api

import (
	"fmt"
	BlockStructs "go-blockchain/blockstructs"
	Commands "go-blockchain/commands"
	Read "go-blockchain/read"
	"log"
	"net/http"
	"net/url"
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

	router.GET("/get_balance", func(c *gin.Context) {
		// Get the private key from the request query parameter
		privateKeyStr, err := url.QueryUnescape(c.Query("private_key"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding private key URL component"})
			return
		}
		if privateKeyStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Private key is required"})
			return
		}

		// Call the GetBalanceByPrivateKey function and return the result
		balance, err := Commands.GetBalanceByPrivateKey(b, privateKeyStr)
		if err != nil {
			fmt.Printf("Error getting balance: %v\n", err) // Add this line to log the error
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"balance": balance,
		})
	})
	//http://localhost:8080/get_balance?private_key=T64iotod7d4re7J/sPC5ZoHqlmJp9TPrGgvok1Npwrc=
	//change PK as needed

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

	// Start the API server
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}

}
