package application

import "omnisys.io/core/domain"

type OrderApp interface {
	Create(order *domain.Order) (*domain.Order, error)
	ViewById(id int) (*domain.Order, error)
}
