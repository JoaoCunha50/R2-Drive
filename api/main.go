package main

import (
	"api/internal/config"
	"api/internal/repositories"
	"api/internal/utils"
	"api/router"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	env := config.EnvVariables
	db := config.CreateDBconnection(env.DATABASE_URL)
	err := config.CreateAdmin(env.ADMIN_USERNAME, env.ADMIN_EMAIL, env.ADMIN_PASSWORD, db)
	utils.InitSecret(env.JWT_SECRET)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Admin created successfully!")
	}

	repos := repositories.InitRepos(db)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))
	r.OPTIONS("/*path", func(c *gin.Context) { c.Status(204) })

	router.SetupMainRouter(r, repos)
	r.Run(":" + env.PORT)
}