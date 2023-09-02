package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Port string
var ConectionDatabase string
var ServerHeader string
var StrictRouting bool
var BodyLimit int
var AzureStorage string
var Container string

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
		//return useError.ErrorEnvironmentVariable
	}

	initVariables()
	return nil
}

func initVariables() {
	Port = os.Getenv("PORT")
	ConectionDatabase = os.Getenv("CONECTION_DATA_BASE")
	ServerHeader = os.Getenv("SERVER_HEADER")
	AzureStorage = os.Getenv("CONECTION_AZURE_STORAGE")
	Container = os.Getenv("CONTAINER_AZURE_STORAGE")
	StrictRouting, _ = strconv.ParseBool(os.Getenv("STRICT_ROUTING"))
	BodyLimit, _ = strconv.Atoi(os.Getenv("BODY_LIMIT"))
}
