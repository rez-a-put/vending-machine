package model

import "time"

type Item struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Price     float64    `json:"price"`
	Status    int8       `json:"status"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
