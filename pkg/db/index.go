package db

import (
	"fmt"
	"os"

	cc "github.com/frankmeza/roomchat/pkg/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbConnection() *gorm.DB {
	postgresDb := postgres.Open(os.Getenv(cc.DB_URL))
	dbConn, err := gorm.Open(postgresDb, &gorm.Config{})

	if err != nil {
		fmt.Println("DB_ERROR IS ", err)
		panic(cc.DB_ERROR)
	}

	return dbConn
}
