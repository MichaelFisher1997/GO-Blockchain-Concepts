package api

import (
	BlockStructs "go-blockchain/blockstructs"
	Commands "go-blockchain/commands"
	Read "go-blockchain/read"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ApiRun(b *BlockStructs.Blockchain) {
	// Set up Gin router
	router := gin.Default()

	// API endpoints
	/*router.POST("/create_wallet", func(c *gin.Context) {
		// Get the initial amount from the request JSON body
		var json struct {
			InitialAmount float64 `json:"initial_amount"`
		}
	
		if err := c.BindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		// Call the MakeWalletWithAmount function and return the result
		privateKeyStr, publicKeyStr, err := Commands.MakeWalletWithAmount(b, json.InitialAmount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{
			"private_key": privateKeyStr,
			"public_key":  publicKeyStr,
		})
		Read.Sync(b)

	})*/
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
	

	// Add other API endpoints for your blockchain commands
	// ...

	// Start the API server
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}

}
