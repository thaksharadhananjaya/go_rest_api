package database

import (
    "log"
    "restapi/internal/domain/user"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func NewPostgresDB(dsn string) *gorm.DB {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database: ", err)
    }

    // Auto migrate the User entity
    if err := db.AutoMigrate(&user.User{}); err != nil {
        log.Fatal("Failed to migrate database: ", err)
    }

    return db
}
