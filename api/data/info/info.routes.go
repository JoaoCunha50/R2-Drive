package info

import (
	"api/data/translations"
	"api/data/users"

	"github.com/gin-gonic/gin"
)

func RegisterInfoRoutes(r *gin.RouterGroup, usersRepo *users.UserRepository, translationsRepo *translations.TranslationRepository) {
	infoService := NewInfoService(usersRepo, translationsRepo)
	infoHandler := NewInfoHandler(infoService)

	r.GET("/", infoHandler.GetInfo)
}