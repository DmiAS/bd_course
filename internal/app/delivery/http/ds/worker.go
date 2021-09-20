package ds

type Worker struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Grade     string `json:"grade"`
	Position  string `json:"position"`
}

type CreateWorkerInput struct {
	Worker
	Login string `json:"login"`
}

type CreateWorkerOutput struct {
	Login    string `json:"login"`
	Password string `json:"gen"`
}

type UpdateWorkerInput struct {
	Worker
}
