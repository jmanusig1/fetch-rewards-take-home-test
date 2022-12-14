package database

import (
	"golang-fetch-api/models"
)

var Balances = make(map[string]int)
var Transactions = make([]models.Transaction, 0)