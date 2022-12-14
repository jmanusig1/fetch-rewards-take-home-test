package managers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"golang-fetch-api/models"
	"golang-fetch-api/database"
	"errors"
)

func TestGetTransactions(t *testing.T) {
	//Arrange
	database.Balances = make(map[string]int)
	database.Transactions = make([]models.Transaction, 0)

	var expectedResult = make([]models.Transaction, 0)

	//Act
	actualResult := GetTransactions()

	//Assert
	assert.Equal(t, expectedResult, actualResult)
}

func TestAddTransaction(t *testing.T) {
	//Arrange
	expectedResult := models.Transaction{"test1", 1, time.Now()}

	//Act
	actualResult, err := AddTransaction(expectedResult)

	//Assert
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, nil, err)
}

func TestBadAddTransaction(t *testing.T) {
	//Arrange
	expectedResult := models.Transaction{"test1", -10, time.Now()}
	expectedError := errors.New("Transaction made payer/user balance negative")

	//Act
	actualResult, actualError := AddTransaction(expectedResult)

	//Assert
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, expectedError, actualError)
}

func TestGetSortedTransactions(t *testing.T) {
	//Arrange
	database.Balances = make(map[string]int)
	database.Transactions = make([]models.Transaction, 0)
	
	transaction1 := models.Transaction{"test1", 1, time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)}
	transaction2 := models.Transaction{"test2", 1, time.Date(2010, time.November, 10, 23, 0, 0, 0, time.UTC)}

	var expectedResult = []models.Transaction{transaction1, transaction2}

	//Act
	AddTransaction(transaction2)
	AddTransaction(transaction1)
	actualResult := GetSortedTransactions()

	//Assert 
	assert.Equal(t, expectedResult, actualResult)
}

func TestSpendPoints(t *testing.T){
	//Arrange
	expectedBalance := models.Balance{"test1", -1}
	var expectedResult = []models.Balance{expectedBalance}

	//Act
	actualResult, actualError := SpendPoints(1)

	//Assert
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, nil, actualError)
}

func TestNegativeSpendPoints(t *testing.T){
	//Arrange
	var expectedResult []models.Balance
	expectedError := errors.New("Cannot spend negative points")

	//Act
	actualResult, actualError := SpendPoints(-1)

	//Assert
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, expectedError, actualError)
}

func TestOverSpendPoints(t *testing.T){
	//Arrange
	var expectedResult []models.Balance
	expectedError := errors.New("Spending more than the allocated points in the user balance")

	//Act
	actualResult, actualError := SpendPoints(2)

	//Assert
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, expectedError, actualError)
}

func TestSpendPointsMultiplePayers(t *testing.T){
	//Arrange
	database.Balances = make(map[string]int)
	database.Transactions = make([]models.Transaction, 0)

	transaction1 := models.Transaction{"test1", 100, time.Date(2011, time.November, 10, 23, 0, 0, 0, time.UTC)}
	transaction2 := models.Transaction{"test2", 100, time.Date(2010, time.November, 10, 23, 0, 0, 0, time.UTC)}
	transaction3 := models.Transaction{"test3", 100, time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)}

	AddTransaction(transaction1)
	AddTransaction(transaction2)
	AddTransaction(transaction3)

	balance1 := models.Balance{"test3", -100}
	balance2 := models.Balance{"test2", -50} 
	expectedResult := []models.Balance{balance1, balance2}

	expectedBalances := map[string]int {
		"test1": 100,
		"test2": 50,
		"test3": 0,
	}

	//Act
	actualResult, actualError := SpendPoints(150)
	actualBalance := GetBalances()

	//Assert
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, nil, actualError)
	assert.Equal(t, expectedBalances, actualBalance)
}

func TestSpendPointsMultiplePayersDifferentTimes(t *testing.T){
	//Arrange
	database.Balances = make(map[string]int)
	database.Transactions = make([]models.Transaction, 0)

	transaction1 := models.Transaction{"test1", 300, time.Date(2000, time.November, 10, 23, 0, 0, 0, time.UTC)}
	transaction2 := models.Transaction{"test2", 200, time.Date(2001, time.November, 10, 23, 0, 0, 0, time.UTC)}
	transaction3 := models.Transaction{"test1", -200, time.Date(2002, time.November, 10, 23, 0, 0, 0, time.UTC)}
	transaction4 := models.Transaction{"test3", 10000, time.Date(2003, time.November, 10, 23, 0, 0, 0, time.UTC)}
	transaction5 := models.Transaction{"test1", 1000, time.Date(2004, time.November, 10, 23, 0, 0, 0, time.UTC)}

	AddTransaction(transaction1)
	AddTransaction(transaction2)
	AddTransaction(transaction3)
	AddTransaction(transaction4)
	AddTransaction(transaction5)

	balance1 := models.Balance{"test1", -100}
	balance2 := models.Balance{"test2", -200} 
	balance3 := models.Balance{"test3", -4700} 
	expectedResult := []models.Balance{balance1, balance2, balance3}

	expectedBalances := map[string]int {
		"test1": 1000,
		"test2": 0,
		"test3": 5300,
	}

	//Act
	actualResult, actualError := SpendPoints(5000)
	actualBalance := GetBalances()

	//Assert
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, nil, actualError)
	assert.Equal(t, expectedBalances, actualBalance)
}