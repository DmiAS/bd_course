package models

type Worker struct {
	User     User `gorm:"embedded"`
	Grade    string
	Position string
}
type Workers []Worker
