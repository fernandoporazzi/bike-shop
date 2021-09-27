package entity

import "time"

type Image struct {
	ID     string `json:"id"`
	BikeId string `json:"bike_id"`
}

type Bike struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Images    []Image   `json:"images"`
	Stock     int32     `json:"stock"`
}
