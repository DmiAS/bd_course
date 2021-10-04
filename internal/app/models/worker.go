package models

type Worker struct {
	User     User   `gorm:"embedded"`
	Grade    string `json:"grade"`
	Position string `json:"position"`
}
type Workers []Worker

type WorkersList struct {
	Amount int `json:"amount"`
	Workers
}

func NewWorkersList(ws Workers) *WorkersList {
	return &WorkersList{
		Amount:  len(ws),
		Workers: ws,
	}
}
