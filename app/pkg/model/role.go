package model

type Role struct {
	Id   int64  `json:"-"`
	Name string `json:"name"`
}
