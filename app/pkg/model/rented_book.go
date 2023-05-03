package model

import "time"

type RentedBook struct {
	Id         int64      `json:"-"`
	RentedOn   time.Time  `json:"rented_on"`
	RentedFor  uint       `json:"rented_for"`
	RentedBook Book       `json:"rented_book"`
	Rentee     User       `json:"rentee"`
	RentStatus RentStatus `json:"rent_status"`
}
