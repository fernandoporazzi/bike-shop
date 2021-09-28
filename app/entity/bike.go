package entity

import "time"

type Bike struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Images    []string  `json:"images"`
	Stock     int32     `json:"stock"`
}
