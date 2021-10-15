package models

type WorkerEntity struct {
	User
	Grade    string `json:"grade"`
	Position string `json:"position"`
}
type Worker struct {
	User     User `gorm:"embedded"`
	Grade    string
	Position string
}
type Workers []Worker

type WorkersList struct {
	Amount  int            `json:"amount"`
	Workers []WorkerEntity `json:"workers"`
}
