package user

import (
	"context"

	"github.com/suvrick/kiss/pkg/db/client/postgres"
)

type UserRepository struct {
	client *postgres.Client
}

func NewUserRepository(db *postgres.Client) *UserRepository {
	return &UserRepository{
		client: db,
	}
}

func (repo *UserRepository) Create(ctx context.Context, user User) (User, error) {

	cmd := `INSERT INTO users(
		email,
		password_hash, 
		at_create,
		at_update
	) 
	VALUES($1, $2, $3, $4)
	RETURNING id`

	row := repo.client.Pool.QueryRow(ctx, cmd,
		&user.Email,
		&user.PasswordHash,
		&user.AtCreate,
		&user.AtUpdate,
	)

	err := row.Scan(&user.ID)

	return user, err
}

func (repo *UserRepository) Update(ctx context.Context, user User) (User, error) {
	return User{}, nil
}

func (repo *UserRepository) Delete(ctx context.Context, userID uint64) error {
	return nil
}

func (repo *UserRepository) Get(ctx context.Context, limit int) ([]User, error) {

	cmd := `
	 SELECT 
	 	id, email, role, at_create, at_update 
	 FROM 
	 	users
	`
	rows, err := repo.client.Pool.Query(ctx, cmd)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Email, &user.Role, &user.AtCreate, &user.AtUpdate)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo *UserRepository) GetByID(ctx context.Context, userID uint64) (User, error) {

	user := User{
		ID: userID,
	}

	cmd := `
	SELECT 
		email, 
		at_create, 
		at_update
	FROM users 
	WHERE id = $1
	LIMIT 1
	`

	err := repo.client.Pool.QueryRow(ctx, cmd, userID).Scan(&user.Email, &user.AtCreate, &user.AtUpdate)

	return user, err
}

func (repo *UserRepository) FindByEmailAndPassword(ctx context.Context, email string, passwordHash string) (User, error) {

	user := User{
		Email:        email,
		PasswordHash: passwordHash,
	}

	cmd := `
			SELECT id, at_create, at_update
			FROM users 
			WHERE email = $1 
			AND password_hash = $2 
			LIMIT 1
	`

	row := repo.client.Pool.QueryRow(ctx, cmd, email, passwordHash)

	err := row.Scan(&user.ID, &user.AtCreate, &user.AtUpdate)

	return user, err
}
