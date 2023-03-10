package user

import (
	"context"
	"fmt"

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
		role,
		password_hash, 
		at_create,
		at_update
	) 
	VALUES($1, $2, $3, $4, $5)
	RETURNING id`

	user.Role = "player"

	row := repo.client.Pool.QueryRow(ctx, cmd,
		&user.Email,
		&user.Role,
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

	cmd := `DELETE FROM users WHERE id = $1`

	tag, err := repo.client.Pool.Exec(ctx, cmd, userID)

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("user not found")
	}

	return err
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
			role,
			at_create, 
			at_update
		FROM users 
		WHERE id = $1
		LIMIT 1
	`

	err := repo.client.Pool.QueryRow(ctx, cmd, userID).Scan(&user.Email, &user.Role, &user.AtCreate, &user.AtUpdate)

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
