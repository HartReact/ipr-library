package model

type Author struct {
	Id   int64  `json:"-"`
	Name string `json:"name"`
	Info string `json:"info"`
}
