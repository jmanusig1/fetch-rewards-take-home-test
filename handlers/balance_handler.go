package handlers

import (
	"golang-fetch-api/managers"
	"github.com/gofiber/fiber/v2"
)

func GetBalances(c *fiber.Ctx) error {
	return c.Status(200).JSON(managers.GetBalances())
}