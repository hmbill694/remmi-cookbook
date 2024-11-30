package model

import "time"

type Recipe struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	EditedAt    time.Time `json:"editedAt"`
}
