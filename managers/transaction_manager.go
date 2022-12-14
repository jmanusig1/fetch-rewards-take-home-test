package managers

import (
	"golang-fetch-api/models"
	"golang-fetch-api/database"
	"time"
	"errors"
	"sort"
)

func AddTransaction(transaction models.Transaction) (models.Transaction, error) {

	ableToAdd := UpdateOrAddBalance(models.Balance{transaction.Payer, transaction.Points})

	if ableToAdd {
		database.Transactions = append(database.Transactions, transaction)
		return transaction, nil
	}

	return transaction, errors.New("Transaction made payer/user balance negative")
}

func GetTransactions() []models.Transaction {
	return database.Transactions
}

func GetSortedTransactions() []models.Transaction {
	sortedTransactions := database.Transactions

	//sort in ascending order by time stamp
	sort.SliceStable(sortedTransactions, func(i, j int) bool {
		return sortedTransactions[i].TimeStamp.Before(sortedTransactions[j].TimeStamp)
	})

	return sortedTransactions
}


func SpendPoints(spendPoints int) ([]models.Balance, error) {

	var spendResult []models.Balance

	if spendPoints < 0 {
		return spendResult, errors.New("Cannot spend negative points")
	}

	if spendPoints == 0 {
		return spendResult, nil
	}

	totalPoints := 0
	for _, points := range database.Balances {
		totalPoints += points
	}

	if spendPoints > totalPoints{
		return spendResult, errors.New("Spending more than the allocated points in the user balance")
	} else  {

		sortedTransactions := GetSortedTransactions()

		var maxNegativePoints = make(map[string]int)
		//first pass to collect the negative balances
		for _, transaction := range sortedTransactions {
			if transaction.Points < 0 {
				maxNegativePoints[transaction.Payer] += transaction.Points
			}
		}

		//then we see how much each payer should pay
		var pointsPaid = make(map[string]int)
		sumPoints := 0

		for _, transaction := range sortedTransactions {
			if transaction.Points > 0 {
				if _, doesMapContainKey := pointsPaid[transaction.Payer]; doesMapContainKey {
					pointsPaid[transaction.Payer] += transaction.Points
					sumPoints += transaction.Points
 				} else {
					pointsPaid[transaction.Payer] += transaction.Points + maxNegativePoints[transaction.Payer]
					sumPoints += transaction.Points + maxNegativePoints[transaction.Payer]
				}
			}

			if sumPoints >= spendPoints {
				pointsPaid[transaction.Payer] -= (sumPoints - spendPoints)
				break
			}
		}
		
		for payer, points := range pointsPaid {
			spendResult = append(spendResult, models.Balance{payer, -points})
			AddTransaction(models.Transaction{payer, -points, time.Now()})
		} 
 	}

	return spendResult, nil
}
