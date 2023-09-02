package utils

import (
	"os"
	"strconv"

	"github.com/c-go-m/go-base/common/useError"
	"github.com/joho/godotenv"
)

var Port string
var ConectionDatabase string
var ServerHeader string
var StrictRouting bool
var BodyLimit int

func LoadEnv() error {
	err := godotenv.Load()

	if err != nil {
		return useError.ErrorEnvironmentVariable
	}

	initVariables()

	return nil
}

func initVariables() {
	Port = os.Getenv("PORT")
	ConectionDatabase = os.Getenv("CONECTION_DATA_BASE")
	ServerHeader = os.Getenv("SERVER_HEADER")
	StrictRouting, _ = strconv.ParseBool(os.Getenv("STRICT_ROUTING"))
	BodyLimit, _ = strconv.Atoi(os.Getenv("BODY_LIMIT"))
}
