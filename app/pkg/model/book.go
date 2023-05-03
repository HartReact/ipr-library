package model

import "time"

type Book struct {
	Id           int64       `json:"-"`
	Name         string      `json:"name"`
	WritingDate  time.Time   `json:"writing_date"`
	PrintingDate time.Time   `json:"printing_date"`
	AddedAt      time.Time   `json:"added_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	Cover        int64       `json:"-"`
	Abstract     string      `json:"abstract"`
	Authors      []Author    `json:"authors"`
	Publishers   []Publisher `json:"publishers"`
	Genres       []Genre     `json:"genres"`
}
