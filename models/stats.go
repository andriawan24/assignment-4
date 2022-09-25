package models

import "time"

type Stats struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Status struct {
	Stats     Stats     `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}
