package contracts

import "product_service_saga/internal/domain"

type ProductRepoContract interface {
	GetProducts() ([]*domain.Product, error)
	GetProductsPaginate(page, pageSize int) ([]*domain.Product, int64, error)
	FindByID(id int) (*domain.Product, error)
}

type ProductServiceContract interface {
	GetProducts() ([]*domain.Product, error)
	GetProductsPaginate(page, pageSize int) ([]*domain.Product, int64, error)
	FindByID(id int) (*domain.Product, error)
}
