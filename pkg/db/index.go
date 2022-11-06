package db

import (
	"os"
	"time"

	cc "github.com/frankmeza/roomchat/pkg/constants"
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
	postgresDb := postgres.Open(os.Getenv(cc.DB_URL))
	dbConn, err := gorm.Open(postgresDb, &gorm.Config{})

	if err != nil {
		return nil, errata.CreateError("GetDbConnection Open", err)
	}

	return dbConn, nil
}
