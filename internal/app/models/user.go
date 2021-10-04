package models

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	VkLink    string    `json:"vk_link"`
	TgLink    string    `json:"tg_link"`
}
type Users []User
