package main

import (
	"context"
	"fmt"
	"go-order-process/datalayer"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4"
)

const webPort = "8091"

type Config struct {
	Repo datalayer.Queries
	Client *http.Client
}

func main() {
	// connect to db
	conn := initDB()
	defer conn.Close(context.Background())
	q := datalayer.New(conn)

	// create channels

	// create waitgroup

	// setup application config
	app := Config{
		Repo: *q,
		Client: &http.Client{},
	}

	// start server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

	// listen for orders
}

func initDB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
    if err != nil {
            fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
            log.Panic("cant connect to db")
    }
	log.Println("Connected to DB !")

	if conn == nil {
		log.Panic("cant connect to db")
	}
	return conn
}