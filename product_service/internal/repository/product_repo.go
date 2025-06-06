package repository

import (
	"log"
	"product_service_saga/internal/contracts"
	"product_service_saga/internal/domain"

	"gorm.io/gorm"
)

type IGorm interface {
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
}

type ProductRepo struct {
	db IGorm
}

func NewProductRepo(db IGorm) contracts.ProductRepoContract {
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
	err := p.db.First(&product, id).Error
	if err != nil {
		log.Printf("error get product by id: %v", err.Error())
		return nil, err
	}
	return product, nil
}
