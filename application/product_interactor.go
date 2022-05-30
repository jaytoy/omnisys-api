package application

import "omnisys.io/core/domain"

type ProductInteractor struct {
	productApp ProductApp
}

// Index is display a listing of the resource.
func (pi *ProductInteractor) Index() (products domain.Products, err error) {
	products, err = pi.productApp.ViewAll()

	return
}

func (pi *ProductInteractor) Show(productID int) (p domain.Product, err error) {
	p, err = pi.productApp.ViewByID(productID)

	return
}

// Store is store a newly created resource in storage.
func (pi *ProductInteractor) Store(p domain.Product) (id int64, err error) {
	id, err = pi.productApp.Create(p)

	return
}

// Destroy is update the specified resource from storage.
func (pi *ProductInteractor) Update(productID int) (err error) {
	err = pi.productApp.EditByID(productID)

	return
}

// Destroy is remove the specified resource from storage.
func (pi *ProductInteractor) Destroy(productID int) (err error) {
	err = pi.productApp.DeleteByID(productID)

	return
}
