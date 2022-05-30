package application

import "omnisys.io/core/domain"

type CustomerApp interface {
	Create(customer *domain.Customer) (*domain.Customer, error)
	ViewAll() (*[]domain.Customer, error)
	ViewById(id int) (*domain.Customer, error)
	Edit(id int, customer *domain.Customer) (*domain.Customer, error)
	Delete(id int) error
}
