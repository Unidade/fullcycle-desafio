package model

import (
	"errors"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

func (p *Product) isValid() error {
	if p.Price <= 0 {
		return errors.New("price product cannot be less or equal to zero")
	}

	return nil
}

func NewProduct(name string, description string, price float32) (*Product, error) {
	newProduct := &Product{
		Name:        name,
		Price:       price,
		Description: description,
	}

	err := newProduct.isValid()
	if err != nil {
		return nil, err
	}

	return newProduct, nil

}

type ProductRepository interface {
	AddProduct(product Product) (*Product, error)
	FindProductByPrice(price float32) (*[]Product, error)
}
