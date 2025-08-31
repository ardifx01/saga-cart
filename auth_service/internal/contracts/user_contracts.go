package contracts

import (
	"auth_service/internal/domain"
	"context"
)

type UserServiceContract interface {
	Create(context.Context, domain.UserCreate) (*domain.User, error)
	FindById(context.Context, int) (*domain.User, error)
	FindByUsername(context.Context, string) (*domain.User, error)
}

type UserRepoContract interface {
	Create(context.Context, domain.UserCreate) (*domain.User, error)
	FindById(context.Context, int) (*domain.User, error)
	FindByUsername(context.Context, string) (*domain.User, error)
}
