package contracts

import "product_service_saga/internal/domain"

type ProductRepoContract interface {
	GetProducts() ([]*domain.Product, error)
	FindByID(id int) (*domain.Product, error)
}

type ProductServiceContract interface {
	GetProducts() ([]*domain.Product, error)
	FindByID(id int) (*domain.Product, error)
}
