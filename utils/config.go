package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var HOST string
var PORT int
var TRUSTED_PROXIES string
var ACCOUNT_SERVICE_GRPC_HOST string
var ACCOUNT_SERVICE_GRPC_PORT int

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var errPort error
	var errAccSerPort error

	HOST = os.Getenv("HOST")
	PORT, errPort = strconv.Atoi(os.Getenv("PORT"))
	ACCOUNT_SERVICE_GRPC_HOST = os.Getenv("ACCOUNT_SERVICE_GRPC_HOST")
	ACCOUNT_SERVICE_GRPC_PORT, errAccSerPort = strconv.Atoi(os.Getenv("ACCOUNT_SERVICE_GRPC_PORT"))
	TRUSTED_PROXIES = os.Getenv("TRUSTED_PROXIES")

	if errPort != nil {
		PORT = 3000
	}

	if errAccSerPort != nil {
		PORT = 50000
	}
}
