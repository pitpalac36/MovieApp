package main

import (
	"backend/repository"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DS string
	DB repository.DatabaseRepo
}

func main() {
	var app application

	flag.StringVar(&app.DS, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &repository.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	log.Println("Starting application on port", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
