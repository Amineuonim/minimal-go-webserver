package core

import (
    "fmt"
    "log"
    "os"
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type DataBase struct {
    DB *gorm.DB
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

    var err error
    t.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
    log.Println("Connected to PostgreSQL successfully!")
}