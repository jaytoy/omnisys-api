package domain

import (
	"time"
)

type Order struct {
	ID         uint         `json:"id"`
	OrderItems []*OrderItem `json:"order_items" gorm:"foreignKey:OrderRefer"`
	Status     string       `json:"status"`
	TotalPrice float64      `json:"total_price"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}

func (Order) TableName() string { return "orders" }

type OrderItem struct {
	Name        string  `json:"name"`
	Qty         float64 `json:"quantity"`
	SinglePrice float64 `json:"single_price"`
	OrderRefer  uint    `json:"order_refer"`
}

func (OrderItem) TableName() string { return "order_item" }
