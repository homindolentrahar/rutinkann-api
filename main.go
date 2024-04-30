package main

import (
	"com.homindolentrahar.rutinkann-api/api"
	"com.homindolentrahar.rutinkann-api/db"
	"com.homindolentrahar.rutinkann-api/helper"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	helper.PanicIfError(envErr)

	validation := validator.New()

	postgresDb := db.NewPostgresStorage()
	database, dbError := postgresDb.Connect()
	helper.PanicIfError(dbError)

	apiService := api.NewChiApiService("localhost:8080", database, validation)

	err := apiService.StartServer()
	helper.PanicIfError(err)
}
