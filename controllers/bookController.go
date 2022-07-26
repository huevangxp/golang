package controllers

import "github.com/gofiber/fiber/v2"

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("get books all in here")
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("get book by id")
}

func AddBook(c *fiber.Ctx) error {
	return c.SendString("add book")
}

func UpdateBook(c *fiber.Ctx) error {
	return c.SendString("update book")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("delete book")
}
