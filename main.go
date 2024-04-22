package main

import (
	"com.homindolentrahar.rutinkann-api/api"
	"com.homindolentrahar.rutinkann-api/data/local"
	"com.homindolentrahar.rutinkann-api/helper"
)

func main() {
	postgresDb := local.NewPostgresStorage()
	db, dbError := postgresDb.Connect()
	helper.PanicIfError(dbError)

	apiService := api.NewChiApiService("localhost:8080", db)

	err := apiService.StartServer()
	helper.PanicIfError(err)
}
