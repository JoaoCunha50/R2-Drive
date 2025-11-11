package router

import (
	"api/internal/repositories"

	"github.com/gin-gonic/gin"
)

func SetupMainRouter(r *gin.Engine, repos *repositories.Repos) {
	webGroup := r.Group("/web") 
	WebRouter(webGroup, repos)
}