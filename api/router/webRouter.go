package router

import (
	"api/data/info"
	"api/data/translations"
	"api/data/users"
	"api/internal/repositories"

	"github.com/gin-gonic/gin"
)

func WebRouter(r *gin.RouterGroup, repos *repositories.Repos) {
    userGroup := r.Group("/users")
	translationsGroup := r.Group("/translations")
	infoGroup := r.Group("/info")

    users.RegisterUserRoutes(userGroup, repos.UsersRepo)
	translations.RegisterTranslationRoutes(translationsGroup, repos.TranslationsRepo)
	info.RegisterInfoRoutes(infoGroup, repos.UsersRepo, repos.TranslationsRepo)
}