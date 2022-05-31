package repository

import (
	"fmt"

	"gorm.io/gorm"
	"omnisys.io/core/application"
	"omnisys.io/core/domain"
)

type OrderRepository struct {
	DB *gorm.DB
}

func CreateOrderRepository(DB *gorm.DB) application.OrderApp {
	return &OrderRepository{DB}
}

func (e *OrderRepository) Create(order *domain.Order) (*domain.Order, error) {
	err := e.DB.Save(&order).Error
	if err != nil {
		fmt.Printf("[OrderRepository.Create] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data")
	}
	return order, nil
}

func (e *OrderRepository) ViewById(id int) (*domain.Order, error) {
	var order = domain.Order{}
	err := e.DB.Table("orders").Where("id = ?", id).First(&order).Error
	if err != nil {
		fmt.Printf("[OrderRepository.ReadById] error execute query %v \n", err)
		return nil, fmt.Errorf("id is not exsis")
	}
	return &order, nil
}
