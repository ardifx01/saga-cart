package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"product_service_saga/internal/cache"
	"product_service_saga/internal/contracts"
	"product_service_saga/internal/domain"

	"gorm.io/gorm"
)

type IGorm interface {
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
}

type ProductRepo struct {
	db    IGorm
	cache *cache.Redis
}

func NewProductRepo(db IGorm, cache *cache.Redis) contracts.ProductRepoContract {
	return &ProductRepo{
		db:    db,
		cache: cache,
	}
}

func (p *ProductRepo) GetProducts() ([]*domain.Product, error) {
	var products []*domain.Product

	ctx := context.Background()

	val, err := p.cache.Get(ctx, "product-list")
	if err == nil {
		json.Unmarshal([]byte(val), &products)
		return products, nil
	}

	err = p.db.Find(&products).Error
	if err != nil {
		log.Printf("Error get all products (repo) : %v", err.Error())
		return nil, err
	}

	jsonData, _ := json.Marshal(products)

	_ = p.cache.Set(ctx, "product-list", jsonData, 1)

	return products, nil
}

func (p *ProductRepo) FindByID(id int) (*domain.Product, error) {
	var product *domain.Product

	ctx := context.Background()
	key := fmt.Sprintf("product-%d", id)

	val, err := p.cache.Get(ctx, key)
	if err == nil {
		json.Unmarshal([]byte(val), &product)
		return product, nil
	}

	err = p.db.First(&product, id).Error
	if err != nil {
		log.Printf("error get product by id: %v", err.Error())
		return nil, err
	}

	jsonData, _ := json.Marshal(product)

	_ = p.cache.Set(ctx, key, jsonData, 1)

	return product, nil
}
