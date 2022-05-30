package repository

import (
	"fmt"

	"gorm.io/gorm"
	"omnisys.io/core/application"
	"omnisys.io/core/domain"
)

type ProductRepository struct {
	DB *gorm.DB
}

func CreateProductRepository(DB *gorm.DB) application.ProductApp {
	return &ProductRepository{DB}
}

func (e *ProductRepository) Create(product *domain.Product) (*domain.Product, error) {
	err := e.DB.Save(&product).Error
	if err != nil {
		fmt.Printf("[ProductRepository.Create] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data")
	}
	return product, nil
}

func (e *ProductRepository) ViewAll() (*[]domain.Product, error) {
	var products []domain.Product
	err := e.DB.Find(&products).Error
	if err != nil {
		fmt.Printf("[ProductRepository.ReadAll] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view all data")
	}
	return &products, nil
}

func (e *ProductRepository) ViewById(id int) (*domain.Product, error) {
	var product = domain.Product{}
	err := e.DB.Table("products").Where("id = ?", id).First(&product).Error
	if err != nil {
		fmt.Printf("[ProductRepository.ReadById] error execute query %v \n", err)
		return nil, fmt.Errorf("id is not exsis")
	}
	return &product, nil
}

func (e *ProductRepository) Edit(id int, product *domain.Product) (*domain.Product, error) {
	var upPerson = domain.Product{}
	err := e.DB.Table("products").Where("id = ?", id).First(&upPerson).Updates(&product).Error
	if err != nil {
		fmt.Printf("[ProductRepository.Update] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update data")
	}
	return &upPerson, nil
}

func (e *ProductRepository) Delete(id int) error {
	var product = domain.Product{}
	err := e.DB.Table("products").Where("id = ?", id).First(&product).Delete(&product).Error
	if err != nil {
		fmt.Printf("[ProductRepository.Delete] error execute query %v \n", err)
		return fmt.Errorf("id is not exsis")
	}
	return nil
}
