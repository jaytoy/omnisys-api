package domain

import "time"

type Order struct {
	ID           uint
	OrderItems   []*OrderItem
	Status       string
	Total        float64
	CreationTime time.Time
	UpdateTime   time.Time
}

type OrderItem struct {
	Name        string
	Qty         float64
	SinglePrice float64
}
