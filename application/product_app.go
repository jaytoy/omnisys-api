package application

import "omnisys.io/core/domain"

type ProductApp interface {
	Create(domain.Product) (int64, error)
	ViewAll() (domain.Products, error)
	ViewByID(int) (domain.Product, error)
	EditByID(int) error
	DeleteByID(int) error
}
