package admin

type CreateDto struct {
	Name     string `json:"name" `
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"Password"`
}
