package db

import (
	"collaborative-task/config"
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var Pool *pgxpool.Pool

func Init() {
	var err error
	Pool, err = pgxpool.Connect(context.Background(), config.DbURL)
	if err != nil {
		log.Fatal("Unable to connect to the database:", err)
	}
	log.Println("Connected to the database.")
}

func Close() {
	Pool.Close()
}
