package models

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Role string

const (
	AdminRole  Role = "admin"
	WorkerRole Role = "worker"
	ClientRole Role = "client"
)

func (r *Role) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	*r = Role(str)
	return nil
}

func (r Role) MarshalJSON() ([]byte, error) {
	str := string(r)
	return json.Marshal(&str)
}

func (r *Role) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.Errorf("Failed to convert %v to string", value)
	}

	*r = Role(str)
	return nil
}

func (r Role) Value() (driver.Value, error) {
	return string(r), nil
}

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	VkLink    string    `json:"vk_link"`
	TgLink    string    `json:"tg_link"`
	Role      Role      `json:"-"`
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
