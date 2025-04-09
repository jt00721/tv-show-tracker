package config

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jt00721/tv-show-tracker/infrastructure"
	"github.com/jt00721/tv-show-tracker/internal/handler"
	"github.com/jt00721/tv-show-tracker/internal/repository"
	"github.com/jt00721/tv-show-tracker/internal/routes"
	"github.com/jt00721/tv-show-tracker/internal/usecase"
)

type App struct {
	Router           *gin.Engine
	ShowHandler      *handler.ShowHandler
	WatchlistHandler *handler.WatchlistHandler
}

func NewApp() *App {
	log.Println("Starting app...")
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Could not load .env file, using system environment variables")
	}

	log.Println("Initialising DB...")
	if err := infrastructure.InitDB(); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}

	watchlistRepository := repository.NewWatchlistRepository(infrastructure.DB)
	client := infrastructure.NewTVMazeClient("https://api.tvmaze.com")
	showUsecase := usecase.NewShowUsecase(client)
	watchlistUsecase := usecase.NewWatchlistUsecase(watchlistRepository)
	showHandler := handler.NewShowHandler(showUsecase)
	watchlistHandler := handler.NewWatchlistHandler(watchlistUsecase)

	router := gin.Default()

	router.Static("/static", "./static")

	routes.SetupRoutes(router, showHandler, watchlistHandler)

	return &App{
		Router:           router,
		ShowHandler:      showHandler,
		WatchlistHandler: watchlistHandler,
	}
}

// Run starts the server
func (app *App) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ip := "0.0.0.0:"
	if env := os.Getenv("ENV"); env == "Dev" || env == "development" {
		ip = ":"
	}

	fmt.Println("Server running on port", port)
	app.Router.Run(ip + port)
}
