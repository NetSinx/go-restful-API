package product

import "github.com/NetSinx/go-restful-API/model/domain"

type ProductService interface {
	Create(product domain.Product) (domain.Product, error)
	Update(product domain.Product, productId int) (domain.Product, error)
	Delete(product domain.Product, productId int) error
	FindById(product domain.Product, productId int) (domain.Product, error)
	FindAll(products []domain.Product) ([]domain.Product, error)
}