package info

import (
	"api/data/users"
	"api/internal/extensions"

	"net/http"

	"github.com/gin-gonic/gin"
)

type InfoHandler struct {
	InfoService *InfoService
}

func NewInfoHandler(infoService *InfoService) *InfoHandler {
	return &InfoHandler{InfoService: infoService}
}

func (h *InfoHandler) GetInfo(c *gin.Context) {
	var user users.UserResponseDTO
	currentUserJson, exists := c.Get("currentUser")
	if exists {
		currentUserMap, ok := currentUserJson.(map[string]any)

		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user context"})
			return
		}

		tempUser, err := h.InfoService.UsersRepo.GetUser(uint(currentUserMap["id"].(float64)))
		if err == nil {
			user = users.UserResponseDTO{ 
				Id: tempUser.ID, 
				Email: tempUser.Email, 
				Name: tempUser.Name, 
				Username: tempUser.Username, 
				Role: tempUser.Role, 
				ProfilePic: tempUser.ProfilePic,
			}
		}
	}

	translations := h.InfoService.TranslationsRepo.GetTranslations()

	var response map[string]any
	if user.Id != 0 {
		response = map[string]any{
			"user": user, 
			"translations": translations, 
			"defaultLanguage": extensions.DefaultLanguage, 
			"languagesSupported": extensions.Languages,
		}
	} else {
		response = map[string]any{
			"translations": translations, 
			"defaultLanguage": extensions.DefaultLanguage, 
			"languagesSupported": extensions.Languages,
		}
	}

	c.JSON(http.StatusOK, response)
}
