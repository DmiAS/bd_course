package models

import (
	"github.com/google/uuid"
)

type Auth struct {
	Login    string
	Password string
	Salt     string
	UserID   uuid.UUID
}
type AuthWithRole struct {
	Auth Auth `gorm:"embedded"`
	Role Role
}

type LogPass struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RoleToken struct {
	Token string `json:"token"`
	Role  Role   `json:"role"`
}
