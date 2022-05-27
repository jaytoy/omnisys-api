package application

import "omnisys.io/core/domain"

type CustomerInteractor struct {
	customerRepository CustomerRepository
}

// func CreatePersonUsecase(customerRepository CustomerRepository) person.PersonUsecase {
// 	return &CustomerInteractor{customerRepository}
// }

func (e *CustomerInteractor) Create(person *domain.Customer) (*domain.Customer, error) {
	return e.customerRepository.Create(person)
}

func (e *CustomerInteractor) ReadAll() (*[]domain.Customer, error) {
	return e.customerRepository.ViewAll()
}

func (e *CustomerInteractor) ReadById(id int) (*domain.Customer, error) {
	return e.customerRepository.ViewById(id)
}

func (e *CustomerInteractor) Update(id int, person *domain.Customer) (*domain.Customer, error) {
	return e.customerRepository.Edit(id, person)
}

func (e *CustomerInteractor) Delete(id int) error {
	return e.customerRepository.Delete(id)
}
