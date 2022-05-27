package interfaces

import (
	"fmt"

	"gorm.io/gorm"
	"omnisys.io/core/domain"
)

type CustomerRepository struct {
	DB *gorm.DB
}

// func CreateCustomerRepository(DB *gorm.DB) application.CustomerRepository {
// 	return &CustomerRepository{DB}
// }

func (e *CustomerRepository) Create(person *domain.Customer) (*domain.Customer, error) {
	err := e.DB.Save(&person).Error
	if err != nil {
		fmt.Printf("[CustomerRepository.Create] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data")
	}
	return person, nil
}

func (e *CustomerRepository) ReadAll() (*[]domain.Customer, error) {
	var persons []domain.Customer
	err := e.DB.Find(&persons).Error
	if err != nil {
		fmt.Printf("[CustomerRepository.ReadAll] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view all data")
	}
	return &persons, nil
}

func (e *CustomerRepository) ReadById(id int) (*domain.Customer, error) {
	var person = domain.Customer{}
	err := e.DB.Table("person").Where("id = ?", id).First(&person).Error
	if err != nil {
		fmt.Printf("[CustomerRepository.ReadById] error execute query %v \n", err)
		return nil, fmt.Errorf("id is not exsis")
	}
	return &person, nil
}

// // probably some problem with gorm
// func (e *CustomerRepository) Update(id int, person *domain.Customer) (*domain.Customer, error) {
// 	var upPerson = domain.Customer{}
// 	err := e.DB.Table("person").Where("id = ?", id).First(&upPerson).Update(&person).Error
// 	if err != nil {
// 		fmt.Printf("[CustomerRepository.Update] error execute query %v \n", err)
// 		return nil, fmt.Errorf("failed update data")
// 	}
// 	return &upPerson, nil
// }

func (e *CustomerRepository) Delete(id int) error {
	var person = domain.Customer{}
	err := e.DB.Table("person").Where("id = ?", id).First(&person).Delete(&person).Error
	if err != nil {
		fmt.Printf("[CustomerRepository.Delete] error execute query %v \n", err)
		return fmt.Errorf("id is not exsis")
	}
	return nil
}
