package application

import "omnisys.io/core/domain"

type CustomerPort interface {
	Create(customer *domain.Customer) (*domain.Customer, error)
	ViewAll() (*[]domain.Customer, error)
	ViewById(id int) (*domain.Customer, error)
	Edit(id int, customer *domain.Customer) (*domain.Customer, error)
	Delete(id int) error
}