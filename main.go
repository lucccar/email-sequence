package main

import (
	"email-sequence/internal/data"
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

	sequenceRepo := data.NewSequenceDataAccess(db)
	sequenceService := service.NewSequenceService(sequenceRepo)
	sequenceHandler := handler.NewSequenceHandler(sequenceService)

	stepRepo := data.NewStepDataAccess(db)
	stepService := service.NewStepService(stepRepo)
	stepHandler := handler.NewStepHandler(stepService)

	r := gin.Default()
	r.POST("/sequences", sequenceHandler.CreateSequence)
	r.PUT("/sequences/:id/steps/:stepId", stepHandler.UpdateStep)
	r.DELETE("/sequences/:id/steps/:stepId", stepHandler.DeleteStep)
	r.PATCH("/sequences/:id/tracking", sequenceHandler.UpdateSequenceTracking)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
