package main

import (
	"email-sequence/internal/data/sequence"
	"email-sequence/internal/handler"
	"email-sequence/internal/service"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sequenceRepo := sequence.NewSequenceDataAccess(db)
	sequenceService := service.NewSequenceService(sequenceRepo)
	sequenceHandler := handler.NewSequenceHandler(sequenceService)

	r := gin.Default()
	r.POST("/sequences", sequenceHandler.CreateSequence)
	// Register other routes

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
