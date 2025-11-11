package users

type Role string

const (
	Admin Role = "admin"
	Member  Role = "member"
)

type User struct {
	ID         uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string  `gorm:"not null" json:"name"`
	Username   string  `gorm:"size:50;uniqueIndex;not null" json:"username"`
	Role       Role    `gorm:"type:enum('admin','member');not null;default:'member'" json:"role"`
	Email      string  `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Password   string  `gorm:"not null" json:"password"`
	ProfilePic *string `gorm:"type:text" json:"profilePic,omitempty"`
}
