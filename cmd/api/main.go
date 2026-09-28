package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kmaxsoul/tabaa-store-api/internal/config"
	"github.com/kmaxsoul/tabaa-store-api/internal/database"
)

func main() {
	var cfg *config.Config
	var err error
	cfg, err = config.LoadConfig()

	if err != nil {
		log.Fatal("failed to load configuration: ", err)
	}

	fmt.Println("TEST URL:", cfg.DatabaseURL)

	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DatabaseURL)

	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	defer pool.Close()

	var router *gin.Engine = gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "API is running well",
			"status":   "success",
			"database": "connected",
		})
	})

	router.Run(":" + cfg.ServerPort)
}
