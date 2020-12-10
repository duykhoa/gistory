package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func migrateDB() {
  dsn := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
    os.Getenv("PG_HOST"),
    os.Getenv("PG_USERNAME"),
    os.Getenv("PG_PASSWORD"),
    os.Getenv("PG_DATABASE"),
    os.Getenv("PG_PORT"),
  )

  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if (err != nil) {
    log.Println("Error when connect to postgres database")
    log.Fatal(err.Error())
  }

  Db = db;
  db.AutoMigrate(&HistoricalData{})
}
