package controllers

import "github.com/gofiber/fiber/v2"

func CreateMember(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "Success create new member",
		"data":    nil,
	})
}
