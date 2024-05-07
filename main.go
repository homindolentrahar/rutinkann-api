package main

import (
	"fmt"
	"net/http"
	"os"

	"com.homindolentrahar.rutinkann-api/helper"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	helper.PanicIfError(envErr)

	servAddress := os.Getenv("SERVER_ADDRESS")
	servPort := os.Getenv("SERVER_PORT")
	address := fmt.Sprintf("%s:%s", servAddress, servPort)

	router := InitializeServer()
	http.ListenAndServe(address, router)
}
