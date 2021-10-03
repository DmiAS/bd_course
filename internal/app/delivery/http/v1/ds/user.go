package ds

// initial information for all types of users
type User struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	VkLink    string `json:"vk_link"`
	TgLink    string `json:"tg_link"`
}

type Auth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
