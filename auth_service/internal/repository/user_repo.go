package repository

import (
	"auth_service/internal/domain"
	"auth_service/util"
	"context"
	"log"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (s *UserRepo) Create(ctx context.Context, req domain.UserCreate) (*domain.User, error) {
	hashedPass, err := util.HashPassword(req.Password)
	if err != nil {
		log.Println("[auth repo] error while hashed password:", err)
		return nil, err
	}
	req.Password = hashedPass
	createdUser := domain.User{
		Username: req.Username,
		Password: hashedPass,
	}
	err = s.db.Create(&createdUser).Error
	if err != nil {
		log.Println("[auth repo] error while create user:", err)
		return nil, err
	}
	return &createdUser, nil
}
func (s *UserRepo) FindById(ctx context.Context, userId int) (*domain.User, error) {
	var findedUser domain.User
	err := s.db.First(&findedUser, userId).Error
	if err != nil {
		log.Println("[auth repo] error while find user by id:", err)
		return nil, err
	}
	return &findedUser, nil
}

func (s *UserRepo) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	var findedUser domain.User
	err := s.db.First(&findedUser, "username = ?", username).Error
	if err != nil {
		log.Println("[auth repo] error while find user by id:", err)
		return nil, err
	}
	return &findedUser, nil
}
