package db

import (
    "fmt"
    "log"
    "os"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() {
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")

    dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
        user, password, dbname, host, port)

    var err error
    DB, err = sqlx.Connect("postgres", dsn)
    if err != nil {
        log.Fatalf("Error opening database: %q", err)
    } else {
        fmt.Println("Successfully connected to the database!")
    }
}

func RunMigrations() {
    migrationFile := "scripts/migrate.sql"

    migrationSQL, err := os.ReadFile(migrationFile)  
    if err != nil {
        log.Fatalf("Error reading migration file: %v", err)
    }

    _, err = DB.Exec(string(migrationSQL))
    if err != nil {
        log.Fatalf("Error running migration: %v", err)
    }

    fmt.Println("Migration completed successfully!")
}
