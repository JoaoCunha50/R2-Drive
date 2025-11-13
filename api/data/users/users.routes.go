package users

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup, repo *UserRepository) {
    handler := NewUserHandler(repo)
    
    r.POST("/login", handler.LoginUser)
    r.GET("/:id", handler.GetUser)
    r.POST("", handler.CreateUser)
}