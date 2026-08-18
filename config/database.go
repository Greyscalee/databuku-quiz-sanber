package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}


func buildConnectionString() string {
	if databaseURL := os.Getenv("DATABASE_URL"); databaseURL != "" {
		return databaseURL
	}
	
	host := getEnv("PGHOST", "localhost")
	port := getEnv("PGPORT", "5432")
	user := getEnv("PGUSER", "postgres")
	password := getEnv("PGPASSWORD", "inlinesk8")
	dbname := getEnv("PGDATABASE", "postgres")

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		host, port, user, password, dbname,
	)
}

func ConnectDB() *sql.DB {
	psqlInfo := buildConnectionString()

	db, err := sql.Open("postgres", psqlInfo)
	if err!= nil {
		panic(err)
	}

	if err = db.Ping(); err !=nil {
		panic(err)
	}

	DB = db
	fmt.Println("Connected")

	if err := runMigration(db); err != nil {
		panic(err)
	}
	return db
}

func runMigration(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS book_library (
				id SERIAL PRIMARY KEY,
				book_name VARCHAR(30) NOT NULL,
				author VARCHAR(20) NOT NULL,
				published INT
			)
			`

	_, err := db.Exec(query)
	return err
}
