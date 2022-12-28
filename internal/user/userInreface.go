package user

import (
	"context"
)

type IUserService interface {
	Get(ctx context.Context, limit int) ([]User, error)
	GetByID(ctx context.Context, userID uint64) (User, error)
	Create(ctx context.Context, email string, password string) (User, error)
	Update(ctx context.Context, user User) (User, error)
	Delete(ctx context.Context, userID uint64) error
	SingIn(ctx context.Context, email string, password string) (string, error)
}

type IUserRepository interface {
	Get(ctx context.Context, limit int) ([]User, error)
	GetByID(ctx context.Context, userID uint64) (User, error)
	Create(ctx context.Context, user User) (User, error)
	Update(ctx context.Context, user User) (User, error)
	Delete(ctx context.Context, userID uint64) error
	FindByEmailAndPassword(ctx context.Context, email string, password string) (User, error)
}
