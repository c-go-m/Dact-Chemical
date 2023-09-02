package database

import (
	"github.com/c-go-m/DC-Admin/common/utility/config"
	"github.com/c-go-m/DC-Admin/common/utility/useError"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBase struct {
	Db *gorm.DB
}

func New() (*DataBase, error) {
	db, err := gorm.Open(postgres.New(getConfig()), &gorm.Config{})

	if err != nil {
		return nil, useError.ErrorConnectionDataBase
	}
	var database = DataBase{
		Db: db,
	}

	errorinit := database.initialDatabase()

	if errorinit != nil {
		return nil, errorinit
	}
	return &database, nil
}

func getConfig() postgres.Config {
	db_config := postgres.Config{
		DSN:                  config.ConectionDatabase,
		PreferSimpleProtocol: true,
	}
	return db_config
}
