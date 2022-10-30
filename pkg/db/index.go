package db

import (
	"os"

	cc "github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/errata"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbConnection() (*gorm.DB, error) {
	postgresDb := postgres.Open(os.Getenv(cc.DB_URL))
	dbConn, err := gorm.Open(postgresDb, &gorm.Config{})

	if err != nil {
		return nil, errata.CreateError(errata.ErrataParams{
			Err:     err,
			ErrFunc: "GetDbConnection gorm.Open",
		})
	}

	return dbConn, nil
}
