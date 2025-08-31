package service

import (
	"auth_service/internal/contracts"
	"auth_service/internal/domain"
	"context"
	"log"
)

type UserService struct {
	userRepo contracts.UserRepoContract
}

func NewUserService(userRepo contracts.UserRepoContract) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Create(ctx context.Context, req domain.UserCreate) (*domain.User, error) {
	user, err := s.userRepo.Create(ctx, req)
	if err != nil {
		log.Println("[auth service] error create user:", err)
		return nil, err
	}
	return user, nil
}
func (s *UserService) FindById(ctx context.Context, userId int) (*domain.User, error) {
	user, err := s.userRepo.FindById(ctx, userId)
	if err != nil {
		log.Println("[auth service] error get user:", err)
		return nil, err
	}
	return user, nil
}

func (s *UserService) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	user, err := s.userRepo.FindByUsername(ctx, username)
	if err != nil {
		log.Println("[auth service] error get user:", err)
		return nil, err
	}
	return user, nil
}
