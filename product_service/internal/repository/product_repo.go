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
	Model(value interface{}) (tx *gorm.DB)
	Count(count *int64) (tx *gorm.DB)
	Limit(limit int) (tx *gorm.DB)
	Offset(offset int) (tx *gorm.DB)
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

func (p *ProductRepo) GetProductsPaginate(limit, offSet int) ([]*domain.Product, int64, error) {
	var products []*domain.Product
	var total int64

	if err := p.db.Model(&domain.Product{}).Count(&total).Error; err != nil {
		log.Println("error get total products paginate:", err.Error())
		return nil, 0, nil
	}

	if err := p.db.Limit(limit).Offset(offSet).Find(&products).Error; err != nil {
		log.Println("error get products paginate:", err.Error())
		return nil, 0, err
	}
	return products, total, nil

}

func (p *ProductRepo) GetProducts() ([]*domain.Product, error) {
	var products []*domain.Product

	ctx := context.Background()

	if p.cache != nil { // ✅ protect against nil
		if val, err := p.cache.Get(ctx, "product-list"); err == nil {
			_ = json.Unmarshal([]byte(val), &products)
			return products, nil
		}
	}

	err := p.db.Find(&products).Error
	if err != nil {
		log.Printf("Error get all products (repo) : %v", err.Error())
		return nil, err
	}

	if p.cache != nil { // ✅ protect against nil
		if jsonData, _ := json.Marshal(products); jsonData != nil {
			_ = p.cache.Set(ctx, "product-list", jsonData, 1)
		}
	}

	return products, nil
}

func (p *ProductRepo) FindByID(id int) (*domain.Product, error) {
	var product *domain.Product

	ctx := context.Background()
	key := fmt.Sprintf("product-%d", id)

	if p.cache != nil {
		val, err := p.cache.Get(ctx, key)
		if err == nil {
			json.Unmarshal([]byte(val), &product)
			return product, nil
		}
	}

	err := p.db.First(&product, id).Error
	if err != nil {
		log.Printf("error get product by id: %v", err.Error())
		return nil, err
	}

	if p.cache != nil {
		jsonData, _ := json.Marshal(product)

		_ = p.cache.Set(ctx, key, jsonData, 1)
	}

	return product, nil
}
