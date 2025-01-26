package services

import (
	"context"
	"log"

	"restaurant-management/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB(cfg *config.Config) {
	var err error
	DB, err = pgxpool.New(context.Background(), cfg.GetDSN())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	log.Println("Database connection established.")
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed.")
	}
}
