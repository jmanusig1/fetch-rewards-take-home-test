package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/gofiber/fiber/v2"
	"net/http/httptest"
	"io/ioutil"
)

func TestGetBalances(t *testing.T) {
	//Arrange
	expectedResultCode := 200
	expectedResultBalance := "{}"

	mockApp := fiber.New()
	mockApp.Get("/get-balances", GetBalances)
	req := httptest.NewRequest("GET", "/get-balances", nil)

	//Act
	actualResult, _ := mockApp.Test(req)
	actualBody, _ := ioutil.ReadAll(actualResult.Body)

	//Assert
	assert.Equal(t, expectedResultCode, actualResult.StatusCode)
	assert.Equal(t, expectedResultBalance, string(actualBody))
}