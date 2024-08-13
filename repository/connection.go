package repository

import (
	"log"
	"lookinlabs-test/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	db *gorm.DB
}

func NewConnection(pgConfig config.PGSQLConfig) *Connection {
	db, err := gorm.Open(postgres.Open(pgConfig.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return &Connection{db: db}
}

func (conn *Connection) DB() *gorm.DB {
	return conn.db
}
