package usecases

import "github.com/Unidade/fullcycle-desafio/domain/model"

type ProductUseCase struct {
	ProductRepository model.ProductRepository
}

func (p *ProductUseCase) CreateProduct(name string, description string, price float32) (*model.Product, error) {
	product, error := model.NewProduct(name, description, price)
	if error != nil {
		return nil, error
	}

	savedProduct, error := p.ProductRepository.AddProduct(*product)
	if error != nil {
		return nil, error
	}

	return savedProduct, nil
}

func (p *ProductUseCase) FindProducts(price float32) (*[]model.Product, error) {
	products, err := p.ProductRepository.FindProductByPrice(price)
	if err != nil {
		return nil, err
	}

	return products, nil
}
