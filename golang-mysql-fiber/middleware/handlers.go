package middleware

import (
	"Golang_Projects/golang-my-fiber/entities"
	"Golang_Projects/golang-my-fiber/storage"
	"Golang_Projects/golang-my-fiber/utils"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func DeletePerson(context *fiber.Ctx) error{
	repository := OpenConnection()

	id := context.Params("id")
	personModel := &entities.Person{}

	if id == "" {
		utils.ReturnError(500,"ID field cant be empty",context)
		return nil
	}

	err := repository.DB.Where("id = ?", id).First(personModel).Error
	if err != nil {
		utils.ReturnError(400,"could not delete the person because dont exist",context)
		return nil
	} else {
		repository.DB.Delete(personModel, id)
	}
	utils.ReturnJson("person deleted successfully", context, id)

	return nil
}

func UpdatePerson(context *fiber.Ctx) error {
	personToUpdate := &entities.Person{} 
	repository := OpenConnection()
	person := &entities.Person{}
	id := context.Params("id")

	if err := context.BodyParser(&person); err != nil {
		utils.ReturnError(422, "request failed", context)
		return nil
	}

	err := repository.DB.Table("person").First(&personToUpdate, id).Error
	if err != nil {
		utils.ReturnError(400,"could not get the person because dont exist",context)
		return nil
	}

	personToUpdate.BirthDate = person.BirthDate
	personToUpdate.FirstName = person.FirstName
	personToUpdate.LastName = person.LastName

	if err := repository.DB.Table("person").Save(&personToUpdate).Error; err != nil {
		utils.ReturnError(400,"could not update the person",context)
		return nil
	}
	
	utils.ReturnJson("updated succesfully", context, personToUpdate)

	return nil
}

func GetPerson(context *fiber.Ctx) error {
	repository := OpenConnection()

	id := context.Params("id")
	personModel := &entities.Person{}

	if id == "" {
		utils.ReturnError(500,"ID field cant be empty",context)
		return nil
	}

	err := repository.DB.Where("id = ?", id).First(personModel).Error
	if err != nil {
		utils.ReturnError(400,"could not get the person because dont exist",context)
		return nil
	}

	utils.ReturnJson("person id fetched successfully", context, personModel)

	return nil
}

func AddPerson(context *fiber.Ctx) error {
	person := &entities.Person{}
	repository := OpenConnection()

	err := context.BodyParser(&person)
	if err != nil {
		utils.ReturnError(422, "request failed", context)
		return nil
	}

	err = repository.DB.Create(&person).Error
	if err != nil {
		utils.ReturnError(400, "could not create person", context)
		return nil
	}

	utils.ReturnJson("person has been added", context, person)

	return nil
}

func GetPeople(context *fiber.Ctx) error{
	repository := OpenConnection()
	personModels := &[]entities.Person{}

	err := repository.DB.Find(personModels).Error
	if err != nil {
		utils.ReturnError(400, "could not get the people", context)
		return err
	}

	utils.ReturnJson("people fetched succesfully", context, personModels)

	return nil
}

func OpenConnection() Repository{
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		DBName: os.Getenv("MY_SQL_DBNAME"),
		Password: os.Getenv("MY_SQL_PASSWORD"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}

	err = entities.MigratePerson(db)

	if err != nil {
		log.Fatal("could not migrate db")
	}

	repository := Repository {
		DB: db,
	}

	return repository
}