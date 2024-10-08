package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"admin-dashboard/initializers" // Replace with the correct path for your project
	"admin-dashboard/models"
)

// CreateTransaction creates a new transaction
func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction

	// Bind incoming JSON to the transaction struct
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the current date for the transaction
	transaction.Date = time.Now()

	// Save transaction to the database
	if err := initializers.DB.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// GetTransaction retrieves a transaction by its ID
func GetTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Transaction

	// Query the database for the transaction by ID
	if err := initializers.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// GetTransactions retrieves all transactions
func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction

	// Fetch all transactions from the database
	if err := initializers.DB.Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

// UpdateTransaction updates an existing transaction
func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Transaction

	// Fetch the transaction by ID
	if err := initializers.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	// Bind updated data from request to the transaction struct
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the updated transaction
	if err := initializers.DB.Save(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update transaction"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// DeleteTransaction deletes a transaction by its ID
func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Transaction

	// Find the transaction by ID
	if err := initializers.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	// Delete the transaction from the database
	if err := initializers.DB.Delete(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}

// GetTransactionsByUser retrieves all transactions for a specific user
func GetTransactionsByUser(c *gin.Context) {
	user := c.Param("user")
	var transactions []models.Transaction

	// Find transactions by user
	if err := initializers.DB.Where("user = ?", user).Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

