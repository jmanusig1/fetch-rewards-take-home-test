package managers

import (
	"golang-fetch-api/models"
	"golang-fetch-api/database"
)

func GetBalances() map[string]int {
	return database.Balances
}

func UpdateOrAddBalance(balance models.Balance) bool{ 
	database.Balances[balance.Payer] += balance.Points

	if database.Balances[balance.Payer] >= 0 {
		return true
	}

	database.Balances[balance.Payer] -= balance.Points
	return false
}