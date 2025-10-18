package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *gin.RouterGroup, db *gorm.DB) {
    repo := NewUserRepository(db)
    handler := NewUserHandler(repo)
    
    r.POST("/login", handler.LoginUser)
    r.GET("/:id", handler.GetUser)
    r.POST("/", handler.CreateUser)
}