package db

import (
	"github.com/frankmeza/roomchat/pkg/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbConnection() *gorm.DB {
	postgresDb := postgres.Open(constants.DB_URL)
	dbConn, err := gorm.Open(postgresDb, &gorm.Config{})

	if err != nil {
		panic(constants.DB_ERROR)
	}

	return dbConn
}
