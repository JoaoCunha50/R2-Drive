package info

import (
	"api/data/translations"
	"api/data/users"
)

type InfoService struct {
	UsersRepo *users.UserRepository
	TranslationsRepo *translations.TranslationRepository
}

func NewInfoService(userRepo *users.UserRepository, translationRepo *translations.TranslationRepository) *InfoService {
	return &InfoService{UsersRepo: userRepo, TranslationsRepo: translationRepo}
}
