package main

import (
	"github.com/c-go-m/DC-Admin/common/utility/config"
	"github.com/c-go-m/DC-Admin/common/utility/useError"
	"github.com/c-go-m/DC-Admin/server"
)

func initEnviroment() {
	err := config.LoadEnv()
	useError.ControlError(err)
}

func initServer() {
	server := server.New()
	err := server.InitServer()
	useError.ControlError(err)
}
