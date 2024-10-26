package main

import (
	"log"

	"github.com/capybara120404/watch-tracker/database"
	"github.com/capybara120404/watch-tracker/handlers"
	"github.com/capybara120404/watch-tracker/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	connecter, err := database.OpenOrCreate("watch_tracker.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer connecter.Close()

	repo := repository.NewSeriesInfoRepository(connecter)
	handler := handlers.NewSeriesHandler(repo)

	router := gin.Default()
	setupRoutes(router, handler)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

func setupRoutes(router *gin.Engine, handler *handlers.SeriesHandler) {
	router.GET("/series", handler.GetAllSeriesHandler)
	router.GET("/series/:id", handler.GetSeriesByIdHandler)
}
