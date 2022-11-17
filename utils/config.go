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

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	HOST = os.Getenv("HOST")
	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	TRUSTED_PROXIES = os.Getenv("TRUSTED_PROXIES")

	if err != nil {
		PORT = 3000
	}
}
