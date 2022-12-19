package product

import (
	"github.com/NetSinx/go-restful-API/model/domain"
	"gorm.io/gorm"
)

type productrepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productrepository {
	return &productrepository{
        DB: db,
    }
}

func(repository *productrepository) Create(product domain.Product) (domain.Product, error) {
	err := repository.DB.Create(&product).Error; if err != nil {
		return product, err
	}

	return product, nil
}

func(repository *productrepository) Update(product domain.Product, productId int) (domain.Product, error) {
	err := repository.DB.Where("id = ?", productId).First(&product).Error; if err != nil {
		return product, err
	}

	repository.DB.Updates(&product)

	return product, nil
}

func (repository *productrepository) Delete(product domain.Product, productId int) error {
	err := repository.DB.Where("id = ?", productId).First(&product).Error; if err != nil {
		return err
	}

	repository.DB.Delete(&product)

	return nil
}

func (repository *productrepository) FindById(product domain.Product, productId int) (domain.Product, error) {
	err := repository.DB.Where("id = ?", productId).First(&product).Error; if err != nil {
		return product, err
	}

	return product, nil
}

func (repository *productrepository) FindAll(products []domain.Product) ([]domain.Product, error) {
	err := repository.DB.Find(&products).Error; if err != nil {
		return products, err
	}

	return products, nil
}