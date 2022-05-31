package application

import "omnisys.io/core/domain"

type OrderInteractor struct {
	orderApp OrderApp
}

func CreateOrderApp(orderApp OrderApp) OrderApp {
	return &OrderInteractor{orderApp}
}

func (oi *OrderInteractor) Create(order *domain.Order) (*domain.Order, error) {
	return oi.orderApp.Create(order)
}

func (oi *OrderInteractor) ViewById(id int) (*domain.Order, error) {
	return oi.orderApp.ViewById(id)
}
