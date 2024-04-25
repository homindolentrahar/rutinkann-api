package main

import (
	"com.homindolentrahar.rutinkann-api/api"
	"com.homindolentrahar.rutinkann-api/db"
	"com.homindolentrahar.rutinkann-api/helper"
)

func main() {
	postgresDb := db.NewPostgresStorage()
	database, dbError := postgresDb.Connect()
	helper.PanicIfError(dbError)

	apiService := api.NewChiApiService("localhost:8080", database)

	err := apiService.StartServer()
	helper.PanicIfError(err)
}
