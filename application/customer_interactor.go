package application

import "omnisys.io/core/domain"

type CustomerInteractor struct {
	customerApp CustomerApp
}

func CreateCustomerApp(customerApp CustomerApp) CustomerApp {
	return &CustomerInteractor{customerApp}
}

func (ci *CustomerInteractor) Create(person *domain.Customer) (*domain.Customer, error) {
	return ci.customerApp.Create(person)
}

func (ci *CustomerInteractor) ViewAll() (*[]domain.Customer, error) {
	return ci.customerApp.ViewAll()
}

func (ci *CustomerInteractor) ViewById(id int) (*domain.Customer, error) {
	return ci.customerApp.ViewById(id)
}

func (ci *CustomerInteractor) Edit(id int, person *domain.Customer) (*domain.Customer, error) {
	return ci.customerApp.Edit(id, person)
}

func (ci *CustomerInteractor) Delete(id int) error {
	return ci.customerApp.Delete(id)
}
