package database

import (
	"golang-fetch-api/models"
)

/*
	Description:
		Balances: a map of all the current payers to points that the user has
		Transactions: a list of all the transactions the user have taken
*/
var Balances = make(map[string]int)
var Transactions = make([]models.Transaction, 0)