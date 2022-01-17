package config

import (
	"fmt"
	"os"
)

type Database struct {
	user     string
	password string
	dbName   string
	address  string
	port     string
}

func newDatabase() *Database {
	return &Database{
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASS"),
		dbName:   os.Getenv("DB_NAME"),
		address:  os.Getenv("DB_ADDRESS"),
		port:     os.Getenv("DB_PORT"),
	}
}
func (c Database) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		c.user, c.password, c.address, c.port, c.dbName,
	)
}
