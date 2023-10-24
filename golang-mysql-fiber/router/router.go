package router

import (
	"Golang_Projects/golang-my-fiber/middleware"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	api := app.Group("api/v1")
	api.Get("/getPerson/:id", middleware.GetPerson)
	api.Get("/getPeople", middleware.GetPeople)
	api.Post("/createPerson", middleware.AddPerson)
	api.Delete("/deletePerson/:id", middleware.DeletePerson)
	api.Put("updatePerson/:id", middleware.UpdatePerson)
}

