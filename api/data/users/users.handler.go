package users

import (
	"net/http"
	"strconv"

	"api/internal/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	repo *UserRepository
}

func NewUserHandler(repo *UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input CreateUserRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	input.Password = string(hash)

	newUser := User {
		Name: input.Name,
		Username: input.Username,
		Email: input.Email,
		Role: Member,
		Password: input.Password,
	}

	err = h.repo.CreateUser(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var request LoginRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err1 := h.repo.LoginUser(request.Email, request.Username, request.Password)
	if err1 != nil || user == nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": err1.Error()})
		return
	}

	token, err2 := utils.SignToken(user.ID, user.Email, string(user.Role), user.Name, user.Username)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err2.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"status": "SERVER_SUCCESS", "success": true, "user": user, "token": token})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	
	user, err := h.repo.GetUser(uint(idInt))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userResponse := UserResponseDTO{
		Id: user.ID,
		Name: user.Name,
		Username: user.Username,
		Email: user.Email,
		Role: user.Role,
		ProfilePic: user.ProfilePic,
	}

	c.JSON(http.StatusOK, gin.H{"data": userResponse})
}