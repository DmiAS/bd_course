package models

type Client struct {
	User User `gorm:"embedded"`
}
type Clients []Client

type ClientsList struct {
	Amount int `json:"amount"`
	Clients
}

func NewClientsList(cs Clients) *ClientsList {
	return &ClientsList{
		Amount:  len(cs),
		Clients: cs,
	}
}
