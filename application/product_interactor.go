package application

import "omnisys.io/core/domain"

type ProductInteractor struct {
	productApp ProductApp
}

func CreateProductApp(productApp ProductApp) ProductApp {
	return &ProductInteractor{productApp}
}

func (pi *ProductInteractor) Create(product *domain.Product) (*domain.Product, error) {
	return pi.productApp.Create(product)
}

func (pi *ProductInteractor) ViewAll() (*[]domain.Product, error) {
	return pi.productApp.ViewAll()
}

func (pi *ProductInteractor) ViewById(id int) (*domain.Product, error) {
	return pi.productApp.ViewById(id)
}

func (pi *ProductInteractor) Edit(id int, product *domain.Product) (*domain.Product, error) {
	return pi.productApp.Edit(id, product)
}

func (pi *ProductInteractor) Delete(id int) error {
	return pi.productApp.Delete(id)
}
