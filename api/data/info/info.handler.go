package info

import "github.com/gin-gonic/gin"

type InfoHandler struct {
	InfoService *InfoService
}

func NewInfoHandler(infoService *InfoService) *InfoHandler {
	return &InfoHandler{InfoService: infoService}
}

func (h *InfoHandler) GetInfo(c *gin.Context) {
	
}
