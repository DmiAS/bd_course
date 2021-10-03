package models

type Client struct {
	User User `gorm:"embedded"`
}
type Clients []Client
