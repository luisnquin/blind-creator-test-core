package repository

import (
	"fmt"

	"go-backend-challenge/environment"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func InitAgenciesDB() *gorm.DB {
	environment.InitializeEnv()

	DbUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=require password=%s TimeZone=America/New_York",
		environment.DbHost,
		environment.DbPort,
		environment.DbUser,
		environment.DbName,
		environment.DbPass,
	)

	PGDB, err := gorm.Open(environment.DbEngine, DbUrl)
	if err != nil {
		fmt.Println("Cannot connect to postgres.go database", err)
	} else {
		fmt.Println("Connected to Postgres!")
	}

	return PGDB
}
