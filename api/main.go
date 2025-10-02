package main

import (
	"api/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	env := config.LoadConfig()

	db := config.DBconnection(env.DATABASE_URL)
	log.Printf("Database connected: %s", db)

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Run(":" + env.PORT)
}