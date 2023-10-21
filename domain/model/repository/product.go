package repository

import (
	"github.com/Unidade/fullcycle-desafio/domain/model"
	"gorm.io/gorm"
)

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func (r ProductRepositoryDb) AddProduct(product model.Product) (*model.Product, error) {
	err := r.Db.Create(product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r ProductRepositoryDb) FindProductByPrice(price float32) (*[]model.Product, error) {
	var products []model.Product
	err := r.Db.Where("price = ?", price).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return &products, nil
}
