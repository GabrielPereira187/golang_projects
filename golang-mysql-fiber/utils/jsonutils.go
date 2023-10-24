package utils

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ReturnError(code int, message string, context *fiber.Ctx) {
	context.Status(code).JSON(
		&fiber.Map{"message": message})
}

func ReturnJson(message string, context *fiber.Ctx, obj interface{}) {
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message" : message,
		"data" :  obj,
	})
}