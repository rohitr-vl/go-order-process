package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

func main() {
	// connect to db
	dbConn := initDB()
	dbConn.Ping()

	// q := db.New(dbConn)
	q := db.New(dbConn)

    author, err := q.GetAuthor(context.Background(), 1)
    if err != nil {
            fmt.Fprintf(os.Stderr, "GetAuthor failed: %v\n", err)
            os.Exit(1)
    }

    fmt.Println(author.Name)

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
/*  db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
*/
	// sqlc
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
    if err != nil {
            fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
            os.Exit(1)
    }
    defer conn.Close(context.Background())
	return conn, nil
}
