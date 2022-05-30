package repository

import (
	"fmt"

	"gorm.io/gorm"
	"omnisys.io/core/application"
	"omnisys.io/core/domain"
)

type CustomerRepository struct {
	DB *gorm.DB
}

func CreateCustomerRepository(DB *gorm.DB) application.CustomerApp {
	return &CustomerRepository{DB}
}

func (e *CustomerRepository) Create(customer *domain.Customer) (*domain.Customer, error) {
	err := e.DB.Save(&customer).Error
	if err != nil {
		fmt.Printf("[CustomerRepository.Create] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data")
	}
	return customer, nil
}

func (e *CustomerRepository) ViewAll() (*[]domain.Customer, error) {
	var customers []domain.Customer
	err := e.DB.Find(&customers).Error
	if err != nil {
		fmt.Printf("[CustomerRepository.ReadAll] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view all data")
	}
	return &customers, nil
}

func (e *CustomerRepository) ViewById(id int) (*domain.Customer, error) {
	var customer = domain.Customer{}
	err := e.DB.Table("customer").Where("id = ?", id).First(&customer).Error
	if err != nil {
		fmt.Printf("[CustomerRepository.ReadById] error execute query %v \n", err)
		return nil, fmt.Errorf("id is not exsis")
	}
	return &customer, nil
}

func (e *CustomerRepository) Edit(id int, customer *domain.Customer) (*domain.Customer, error) {
	var upPerson = domain.Customer{}
	err := e.DB.Table("customer").Where("id = ?", id).First(&upPerson).Updates(&customer).Error
	if err != nil {
		fmt.Printf("[CustomerRepository.Update] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update data")
	}
	return &upPerson, nil
}

func (e *CustomerRepository) Delete(id int) error {
	var customer = domain.Customer{}
	err := e.DB.Table("customer").Where("id = ?", id).First(&customer).Delete(&customer).Error
	if err != nil {
		fmt.Printf("[CustomerRepository.Delete] error execute query %v \n", err)
		return fmt.Errorf("id is not exsis")
	}
	return nil
}
