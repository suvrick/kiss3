package user

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/suvrick/kiss/internal/jwthelper"
)

type UserService struct {
	repo IUserRepository
}

func NewUserService(repo IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (srv *UserService) Get(ctx context.Context, limit int) ([]User, error) {
	return srv.repo.Get(ctx, limit)
}

func (srv *UserService) GetByID(ctx context.Context, proxyID uint64) (User, error) {
	return srv.repo.GetByID(ctx, proxyID)
}

func (srv *UserService) Create(ctx context.Context, email, password string) (User, error) {
	user := User{
		Email:        email,
		PasswordHash: getHash(password),
		AtCreate:     time.Now(),
		AtUpdate:     time.Now(),
	}

	return srv.repo.Create(ctx, user)
}

func (srv *UserService) Update(ctx context.Context, user User) (User, error) {
	user.AtUpdate = time.Now()
	return srv.repo.Update(ctx, user)
}

func (srv *UserService) Delete(ctx context.Context, userID uint64) error {
	return srv.repo.Delete(ctx, userID)
}

func (srv *UserService) SingIn(ctx context.Context, email string, password string) (string, error) {

	user, err := srv.repo.FindByEmailAndPassword(ctx, email, getHash(password))
	if err != nil {
		return "", fmt.Errorf("invalid email or password.")
	}

	return jwthelper.NewToken(user.ID, user.Email, user.Role)
}

func getHash(password string) string {
	hasher := sha1.New()
	hasher.Write([]byte(password))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
