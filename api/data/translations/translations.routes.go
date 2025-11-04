package translations

import (
	"github.com/gin-gonic/gin"
)

func RegisterTranslationRoutes(r *gin.RouterGroup, repo *TranslationRepository) {
	handler := NewTranslationHandler(repo)
	
	r.GET("/", handler.GetTranslations)
	r.POST("/", handler.CreateTranslation)
}