package entity

type User struct {
	Id       int64  `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Token string `json:"token,omitempty"`
	Error string `json:"error,omitempty"`
}
