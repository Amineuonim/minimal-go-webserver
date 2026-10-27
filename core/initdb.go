package core

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type DataBase struct {
	DB *sqlx.DB
}

func (t *DataBase) ConnectPGDB() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_TZ"),
	)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	t.DB = db
	log.Println("Connected to PostgreSQL successfully!")
}
