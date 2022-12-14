package handlers

import (
	"golang-fetch-api/models"
	"golang-fetch-api/managers"
	"github.com/gofiber/fiber/v2"
)

/*
	Description: AddTransaction adds a transaction by calling the transaction manager
	Params: takes an fiber context instance
	Return: a HTTP status code along with a JSON return body with the original transaction or error
*/
func AddTransaction(c *fiber.Ctx) error {

	var transaction models.Transaction
	
	//need to add that it is necessary to provide a payer point and timestamp
	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	_, err := managers.AddTransaction(transaction)
	if err != nil { 
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(transaction)
}

/*
	Description: SpendPoints uses user points by calling the transaction manager
	Params: takes an fiber context instance
	Return: a HTTP status code along with a JSON return body with the original transaction or error
*/
func SpendPoints(c *fiber.Ctx) error {
	var points models.Balance
	
	//need to add that it is necessary to provide a payer point and timestamp
	if err := c.BodyParser(&points); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	body, err := managers.SpendPoints(points.Points)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(body)	
}