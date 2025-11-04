package repositories

import (
	"api/data/translations"
	"api/data/users"

	"gorm.io/gorm"
)

type Repos struct {
	UsersRepo *users.UserRepository
	TranslationsRepo *translations.TranslationRepository
}

func InitRepos(db *gorm.DB) *Repos{
	usersRepo := users.NewUserRepository(db)
	translationsRepo := translations.NewTranslationRepository(db)

	return &Repos{UsersRepo: usersRepo, TranslationsRepo: translationsRepo}
}