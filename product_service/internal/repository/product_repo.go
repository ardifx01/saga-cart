package repository

import (
	"log"
	"product_service_saga/internal/contracts"
	"product_service_saga/internal/domain"

	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) contracts.ProductRepoContract {
	return &ProductRepo{
		db: db,
	}
}

func (p *ProductRepo) GetProducts() ([]*domain.Product, error) {
	var products []*domain.Product
	err := p.db.Find(&products).Error
	if err != nil {
		log.Printf("Error get all products (repo) : %v", err.Error())
		return nil, err
	}
	return products, nil
}

func (p *ProductRepo) FindByID(id int) (*domain.Product, error) {
	var product *domain.Product
	err := p.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		log.Printf("error get product by id: %v", err.Error())
		return nil, err
	}
	return product, nil
}
