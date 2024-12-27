package service

import (
	"log"
	"product_service_saga/internal/contracts"
	"product_service_saga/internal/domain"
)

type ProductService struct {
	productRepo contracts.ProductRepoContract
}

func NewProductService(productRepo contracts.ProductRepoContract) contracts.ProductServiceContract {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (s *ProductService) GetProducts() ([]*domain.Product, error) {
	listProducts, err := s.productRepo.GetProducts()
	if err != nil {
		log.Printf("Error get all products (service) : %v", err.Error())
		return nil, err
	}
	return listProducts, nil
}

func (s *ProductService) FindByID(id int) (*domain.Product, error) {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		log.Printf("error get product by id: %v", err.Error())
		return nil, err
	}
	return product, err
}
