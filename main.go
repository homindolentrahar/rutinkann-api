package main

import (
	"net/http"

	"com.homindolentrahar.rutinkann-api/helper"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	helper.PanicIfError(envErr)

	// servAddress := os.Getenv("SERVER_ADDRESS")
	// servPort := os.Getenv("SERVER_PORT")
	// address := fmt.Sprintf("%s:%s", servAddress, servPort)

	router := InitializeServer()
	http.ListenAndServe(":8080", router)
}
