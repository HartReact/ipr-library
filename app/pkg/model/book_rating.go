package model

import "time"

type BookRating struct {
	Id      int64     `json:"-"`
	Rating  float32   `json:"rating"`
	RatedOn time.Time `json:"rated_on"`
	Review  string    `json:"review"`
	Book    Book      `json:"book"`
	User    User      `json:"user"`
}
