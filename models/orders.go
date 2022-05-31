package models

import "time"

type Order struct {
	ID           uint      `gorm:"primaryKey"`
	CustomerName string    `json:"customerName" gorm:"not null;type:varchar(191)"`
	Item         []Item    `json:"Item"`
	OrderedAt    time.Time `json:"orderedAt"`
}

type CreateOrder struct {
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Item         []Item    `json:"Item"`
}
