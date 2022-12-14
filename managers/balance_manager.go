package managers

import (
	"golang-fetch-api/models"
	"golang-fetch-api/database"
)

/*
	Description: GetBalances gets the current balances of the user
	Return: a string-int map containing all the balances of the user
*/
func GetBalances() map[string]int {
	return database.Balances
}

/*
	Description: UpdateOrAddBalance gets the current balances of the user
	Params: takes an instance of the model `balance`
	Return: a boolean indicating whether the action has succeeded
*/
func UpdateOrAddBalance(balance models.Balance) bool{ 
	database.Balances[balance.Payer] += balance.Points

	if database.Balances[balance.Payer] >= 0 {
		return true
	}

	database.Balances[balance.Payer] -= balance.Points
	return false
}