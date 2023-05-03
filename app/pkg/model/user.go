package model

type User struct {
	Id       int64  `json:"-"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Roles    []Role `json:"roles"`
}
