package helpers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DBParams struct {
	dbHost     string
	dbPort     string
	dbUsername string
	dbName     string
	dbPassword string
}

func NewDBParams() *DBParams {
	return &DBParams{
		dbHost:     "localhost",
		dbPort:     "5432",
		dbUsername: "postgres",
		dbName:     "interview",
		dbPassword: "postgres",
	}
}

func NewDB(param *DBParams) *gorm.DB {
	dsn := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable TimeZone=Asia/Bangkok",
		param.dbHost,
		param.dbPort,
		param.dbUsername,
		param.dbName,
		param.dbPassword,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Fail to load database")
	}

	return db
}
