package domain

import (
	"time"
)

type Order struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
	Status     string      `json:"status"`
	TotalPrice float64     `json:"total_price"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

func (Order) TableName() string { return "orders" }

type OrderItem struct {
	LineID      uint    `json:"line_id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Qty         uint    `json:"quantity"`
	SinglePrice float64 `json:"single_price"`
	OrderID     uint    `json:"order_id"`
}

func (OrderItem) TableName() string { return "order_item" }
