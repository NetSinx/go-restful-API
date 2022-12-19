package product

import (
	"errors"

	"github.com/NetSinx/go-restful-API/model/domain"
	repository "github.com/NetSinx/go-restful-API/repository/product"
)

type productservice struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(productService repository.ProductRepository) *productservice {
	return &productservice{
        ProductRepository: productService,
    }
}

func (service *productservice) Create(product domain.Product) (domain.Product, error) {
	createProduct, err := service.ProductRepository.Create(product)
	if err != nil {
		return product, errors.New("failed to create product")
	}

	return createProduct, nil
}

func (service *productservice) Update(product domain.Product, productId int) (domain.Product, error) {
	updateProduct, err := service.ProductRepository.Update(product, productId)
    if err != nil {
		return product, errors.New("the data you want to update is not found")
    }

	return updateProduct, nil
}

func (service *productservice) Delete(product domain.Product, productId int) error {
	err := service.ProductRepository.Delete(product, productId); if err != nil {
		return errors.New("the data you want to delete is not found")
    }

	return nil
}

func (service *productservice) FindById(product domain.Product, productId int) (domain.Product, error) {
	findProduct, err := service.ProductRepository.FindById(product, productId)
    if err != nil {
		err = errors.New("the data you want to find is not found")
		return product, err
    }

	return findProduct, nil
}

func (service *productservice) FindAll(products []domain.Product) ([]domain.Product, error) {
	findProducts, err := service.ProductRepository.FindAll(products)
    if err != nil {
        return findProducts, errors.New("all data not found")
    }

	return findProducts, nil
}