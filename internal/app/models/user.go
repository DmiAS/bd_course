package models

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	VkLink    string    `json:"vk_link"`
	TgLink    string    `json:"tg_link"`
	Role      string
}
type Users []User

type UserList struct {
	Amount int
	Users  Users
}

type UserInfo struct {
	ID   uuid.UUID
	Role Role
}
