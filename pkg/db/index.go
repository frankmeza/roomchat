package db

import (
	"os"
	"time"

	"github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/errata"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	ChildRecord struct {
		CreatedAt time.Time `json:"created_at"`
	}

	DbRecord struct {
		ID        int            `gorm:"primaryKey;autoIncrement" json:"id"`
		CreatedAt time.Time      `json:"created_at"`
		UpdatedAt time.Time      `json:"updated_at"`
		DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	}
)

func GetDbConnection() (*gorm.DB, error) {
	postgresDb := postgres.Open(os.Getenv(constants.DB_URL))
	dbConn, err := gorm.Open(postgresDb, &gorm.Config{})

	if err != nil {
		return nil, errata.CreateError(err, []string{
			"GetDbConnection Open",
		})
	}

	return dbConn, nil
}
