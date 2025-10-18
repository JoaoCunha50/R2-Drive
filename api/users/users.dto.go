package users

type UserResponseDTO struct {
	Id    int `json:"id"`
	Name  string `json:"name"` 
	Username string `json:"username"`
	Email string `json:"email"`
	ProfilePic string `json:"profilePic"`
}

type LoginRequest struct {
    Email    *string `json:"email"`
    Username *string `json:"username"`
    Password string `json:"password"`
}