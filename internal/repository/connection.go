package repository

import (
	"fmt"

	"github.com/luisnquin/blind-creator/test-core/environment"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitAgenciesDB() *gorm.DB {
	environment.InitializeEnv()

	environment.DbPass = "5B$Ns5X2N$PF8GQK9fAr8ueAs3r"

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=require password=%s TimeZone=America/New_York",
		environment.DbHost,
		environment.DbPort,
		environment.DbUser,
		environment.DbName,
		environment.DbPass,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Println("Cannot connect to postgres.go database", err)
	} else {
		fmt.Println("Connected to Postgres!")
	}

	return db
}
