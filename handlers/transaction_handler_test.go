package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/gofiber/fiber/v2"
	"net/http/httptest"
	"io/ioutil"
	"bytes"
)

func TestAddTransaction(t *testing.T) {
	//Arrange
	expectedResultCode := 200
	expectedResultBalance := `{"payer":"test","points":1,"timestamp":"2019-12-12T03:20:50.52-04:00"}`

	mockApp := fiber.New()
	mockApp.Post("/add-transaction", AddTransaction)

	transaction := []byte(`{"payer": "test", "points": 1, "timestamp": "2019-12-12T03:20:50.52-04:00"}`)
	req := httptest.NewRequest("POST", "/add-transaction", bytes.NewBuffer(transaction))
	req.Header.Set("Content-Type", "application/json")

	//Act
	actualResult, _ := mockApp.Test(req)
	actualBody, _ := ioutil.ReadAll(actualResult.Body)

	//Assert
	assert.Equal(t, expectedResultCode, actualResult.StatusCode)
	assert.Equal(t, expectedResultBalance, string(actualBody))
}

func TestBadRequestAddTransaction(t *testing.T) {
	//Arrange
	expectedResultCode := 400

	mockApp := fiber.New()
	mockApp.Post("/add-transaction", AddTransaction)

	transaction := []byte(`random`)
	req := httptest.NewRequest("POST", "/add-transaction", bytes.NewBuffer(transaction))
	req.Header.Set("Content-Type", "application/json")

	//Act
	actualResult, _ := mockApp.Test(req)
	//Assert
	assert.Equal(t, expectedResultCode, actualResult.StatusCode)
}

func TestBadAddTransactionNegativeBalance(t *testing.T) {
	//Arrange
	expectedResultCode := 400
	expectedResultBalance := `"Transaction made payer/user balance negative"`

	mockApp := fiber.New()
	mockApp.Post("/add-transaction", AddTransaction)

	transaction := []byte(`{"payer": "test", "points": -51, "timestamp": "2019-12-12T03:20:50.52-04:00"}`)
	req := httptest.NewRequest("POST", "/add-transaction", bytes.NewBuffer(transaction))
	req.Header.Set("Content-Type", "application/json")

	//Act
	actualResult, _ := mockApp.Test(req)
	actualBody, _ := ioutil.ReadAll(actualResult.Body)

	//Assert
	assert.Equal(t, expectedResultCode, actualResult.StatusCode)
	assert.Equal(t, expectedResultBalance, string(actualBody))
}

func TestSpendPoints(t *testing.T) {
	//Arrange
	expectedResultCode := 200
	expectedResultBalance := "null"

	mockApp := fiber.New()
	mockApp.Post("/spend-points", SpendPoints)

	transaction := []byte(`{"payer": "test", "points": 0}`)
	req := httptest.NewRequest("POST", "/spend-points", bytes.NewBuffer(transaction))
	req.Header.Set("Content-Type", "application/json")

	//Act
	actualResult, _ := mockApp.Test(req)
	actualBody, _ := ioutil.ReadAll(actualResult.Body)

	//Assert
	assert.Equal(t, expectedResultCode, actualResult.StatusCode)
	assert.Equal(t, expectedResultBalance, string(actualBody))
}

func TestBadRequestSpendPoints(t *testing.T) {
	//Arrange
	expectedResultCode := 400

	mockApp := fiber.New()
	mockApp.Post("/spend-points", SpendPoints)

	transaction := []byte(`{"payer": "test", "points": "10"}`)
	req := httptest.NewRequest("POST", "/spend-points", bytes.NewBuffer(transaction))
	req.Header.Set("Content-Type", "application/json")

	//Act
	actualResult, _ := mockApp.Test(req)

	//Assert
	assert.Equal(t, expectedResultCode, actualResult.StatusCode)
}

func TestBadSpendPointsNegativeBalance(t *testing.T) {
	//Arrange
	expectedResultCode := 400
	expectedResultBalance := `"Spending more than the allocated points in the user balance"`

	mockApp := fiber.New()
	mockApp.Post("/spend-points", SpendPoints)

	transaction := []byte(`{"payer": "test", "points": 10}`)
	req := httptest.NewRequest("POST", "/spend-points", bytes.NewBuffer(transaction))
	req.Header.Set("Content-Type", "application/json")

	//Act
	actualResult, _ := mockApp.Test(req)
	actualBody, _ := ioutil.ReadAll(actualResult.Body)

	//Assert
	assert.Equal(t, expectedResultCode, actualResult.StatusCode)
	assert.Equal(t, expectedResultBalance, string(actualBody))
}