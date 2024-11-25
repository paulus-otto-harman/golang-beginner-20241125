package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"project/config"
)

func PgConnect(dbName string, dbUser string, dbPassword string, dbHost string) (*sql.DB, error) {
	return sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", dbUser, dbPassword, dbName, dbHost))
}

func Open(database config.DatabaseConfig) *sql.DB {
	db, err := PgConnect(database.Name, database.Username, database.Password, database.Host)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
