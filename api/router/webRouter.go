package router

import (
	"api/data/info"
	"api/data/translations"
	"api/data/users"

	"github.com/gin-gonic/gin"
)

func WebRouter(r *gin.RouterGroup, usersRepo *users.UserRepository, translationsRepo *translations.TranslationRepository) {
    userGroup := r.Group("/users")
	translationsGroup := r.Group("/translations")
	infoGroup := r.Group("/info")

    users.RegisterUserRoutes(userGroup, usersRepo)
	translations.RegisterTranslationRoutes(translationsGroup, translationsRepo)
	info.RegisterInfoRoutes(infoGroup, usersRepo, translationsRepo)
}