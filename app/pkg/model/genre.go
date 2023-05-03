package model

type Genre struct {
	Id          int64  `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
