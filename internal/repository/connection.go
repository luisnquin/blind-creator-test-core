package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"go-backend-challenge/environment"
)

func InitAgenciesDB() *gorm.DB {
	environment.InitializeEnv()
	var err error
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
