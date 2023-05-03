package model

type BookCollection struct {
	Id          int64  `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Books       []Book `json:"books"`
}
