package info

import "api/data/users"

type translationDetails struct {
	Tag string `json:"tag"`
	Lang string `json:"lang"`
	Translation string `json:"translation"`
}

type InfoDTO struct {
	Languages []string `json:"languages"`
	Translations []translationDetails `json:"translations"`
	User *users.UserResponseDTO `json:"user"`
}