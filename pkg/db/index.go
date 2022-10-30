package db

import (
	"os"

	cc "github.com/frankmeza/roomchat/pkg/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbConnection() (*gorm.DB, error) {
	postgresDb := postgres.Open(os.Getenv(cc.DB_URL))
	dbConn, err := gorm.Open(postgresDb, &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
