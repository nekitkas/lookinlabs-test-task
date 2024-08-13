package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type PGSQLConfig interface {
	DSN() string
}

type pgsqlConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func NewPGSQLConfig() (PGSQLConfig, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	return &pgsqlConfig{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbname:   dbname,
	}, nil
}

func (p pgsqlConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", p.host, p.port, p.user, p.password, p.dbname)
}
