package application

import "omnisys.io/core/domain"

type ProductApp interface {
	Create(product *domain.Product) (*domain.Product, error)
	ViewAll() (*[]domain.Product, error)
	ViewById(id int) (*domain.Product, error)
	Edit(id int, product *domain.Product) (*domain.Product, error)
	Delete(id int) error
}
