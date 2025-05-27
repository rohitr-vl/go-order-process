package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	// connect to db
	db := initDB()
	db.Ping()

	// create channels

	// create waitgroup

	// setup application config

	// start server

	// listen for orders
}

func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("cant connect to db")
	}
	return conn
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	for counts := 0; counts < 10; counts++ {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("postgres not ready yet, wait a sec...")
			time.Sleep(1 * time.Second)
		} else {
			log.Println("connected to db")
			return connection
		}
	}
	log.Println("Unable to connect to db after 10 retries.")
	return nil
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
