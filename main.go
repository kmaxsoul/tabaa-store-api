package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kmaxsoul/tabaa-store-api/config"
	"github.com/kmaxsoul/tabaa-store-api/database"
	"github.com/kmaxsoul/tabaa-store-api/handlers"
	"github.com/kmaxsoul/tabaa-store-api/repository"
)

func main() {
	var cfg *config.Config
	var err error
	cfg, err = config.LoadConfig()

	if err != nil {
		log.Fatal("failed to load configuration: ", err)
	}

	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DatabaseURL)

	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	defer pool.Close()

	userRepo := repository.NewUserRepository(pool)

	userHandler := handlers.NewUserHandler(userRepo)

	var router *gin.Engine = gin.Default()
	router.SetTrustedProxies(nil)

	api := router.Group("/api")
	{
		api.POST("/register", userHandler.RegisterUser)
	}

	router.Run(":" + cfg.ServerPort)
}
