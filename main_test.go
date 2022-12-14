package main 

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"bytes"
    "encoding/json"
	"testing"
	"net/http"
	"io/ioutil"
	"golang-fetch-api/models"
	"time"
)

func SetUpApp() *fiber.App {
	app := fiber.New()
	setupRoutes(app)
	return app
}

func TestHomePage(t *testing.T) {
	//Arrange
	app := SetUpApp()
	req, _ := http.NewRequest("GET", "/", nil)

	//Act
	resp, _ := app.Test(req, 1)
	respBody, _ := ioutil.ReadAll(resp.Body)

	//Assert
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.Equal(t, "Fetch Points Api", string(respBody))
}

func TestGetBalances(t *testing.T) {
	//Arrange
	app := SetUpApp()
	req, _ := http.NewRequest("GET", "/get-balances", nil)

	//Act
	resp, _ := app.Test(req, 1)

	//Assert
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestAddTransaction(t *testing.T) {
	//Arrange
	app := SetUpApp()

	transaction := models.Transaction{
		Payer: "tester", 
		Points: 1, 
		TimeStamp: time.Now(),
	}
	jsonValue, _ := json.Marshal(transaction)

	req, _ := http.NewRequest("POST", "/add-transaction", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	//Act
	resp, _ := app.Test(req, 1)

	//Assert
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestSpendPoints(t *testing.T) {
	//Arrange
	app := SetUpApp()

	points := models.Balance{
		Payer: "tester", 
		Points: 1, 
	}

	jsonValue, _ := json.Marshal(points)
	req, _ := http.NewRequest("POST", "/spend-points", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	//Act
	resp, _ := app.Test(req, 1)

	//Assert
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}