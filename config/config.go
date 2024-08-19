package config

import (
	"fmt"
	"os"
)

type Config struct {
    DatabaseDSN   string
    ServerAddress string
}

func LoadConfig() Config {
	dbUser := os.Getenv("DATABASE_USER")
    dbPassword := os.Getenv("DATABASE_PASSWORD")
    dbName := os.Getenv("DATABASE_NAME")
    dbHost := os.Getenv("DATABASE_HOST")
    dbPort := os.Getenv("DATABASE_PORT")
    port := os.Getenv("PORT",)

    databaseDSN := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s",
        dbUser, dbPassword, dbName, dbHost, dbPort)
    return Config{
        DatabaseDSN: databaseDSN,
        ServerAddress: fmt.Sprintf(":%s",port),
    }
}
