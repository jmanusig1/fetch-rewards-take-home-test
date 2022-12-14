package managers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"golang-fetch-api/models"
	"golang-fetch-api/database"
)

func ResetDatabase() {
	database.Balances = make(map[string]int)
	database.Transactions = make([]models.Transaction, 0)
}

func TestGetBalances(t *testing.T) {
	//Arrange
	ResetDatabase()
	expectedResult := make(map[string]int)

	//Act
	actualResult := GetBalances()

	//Assert
	assert.Equal(t, expectedResult, actualResult)
}

func TestUpdateBalances(t *testing.T) {
	//Arrange
	expectedResult := map[string]int {
		"test1": 1, 
		"test2": 2,
	}

	//Act
	UpdateOrAddBalance(models.Balance{"test1", 1})
	UpdateOrAddBalance(models.Balance{"test2", 2})

	actualResult := GetBalances()

	//Assert
	assert.Equal(t, expectedResult, actualResult)
}