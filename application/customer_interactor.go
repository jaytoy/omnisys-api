package application

import "omnisys.io/core/domain"

type CustomerInteractor struct {
	customerPort CustomerPort
}

// func CreateCustomerApp(customerPort CustomerRepository) CustomerRepository {
// 	return &CustomerInteractor{customerPort}
// }

func (e *CustomerInteractor) Create(person *domain.Customer) (*domain.Customer, error) {
	return e.customerPort.Create(person)
}

func (e *CustomerInteractor) ReadAll() (*[]domain.Customer, error) {
	return e.customerPort.ViewAll()
}

func (e *CustomerInteractor) ReadById(id int) (*domain.Customer, error) {
	return e.customerPort.ViewById(id)
}

func (e *CustomerInteractor) Update(id int, person *domain.Customer) (*domain.Customer, error) {
	return e.customerPort.Edit(id, person)
}

func (e *CustomerInteractor) Delete(id int) error {
	return e.customerPort.Delete(id)
}
