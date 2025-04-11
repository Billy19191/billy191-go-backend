package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	PostgresHost     string
	PostgresPort     string
	PostgresDB       string
	PostgresUsername string
	PostgresPassword string
	DB               *gorm.DB
}

func NewDatabase(host, port, db, username, password string) *Database {
	return &Database{
		PostgresHost:     host,
		PostgresPort:     port,
		PostgresDB:       db,
		PostgresUsername: username,
		PostgresPassword: password,
		DB:               &gorm.DB{},
	}
}

func (d *Database) Connect() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		d.PostgresHost, d.PostgresUsername, d.PostgresPassword, d.PostgresDB, d.PostgresPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	d.DB = db
	log.Println("Database connected successfully")
	return nil
}
