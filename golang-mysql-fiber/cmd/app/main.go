package main

import (
	"Golang_Projects/golang-my-fiber/router"
	"github.com/gofiber/fiber/v2"
)

func main(){

	app := fiber.New()

	router.Router(app)

	app.Listen(":8080")
}