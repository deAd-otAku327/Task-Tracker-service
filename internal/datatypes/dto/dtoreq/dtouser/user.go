package dtouser

type PostLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostRegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
