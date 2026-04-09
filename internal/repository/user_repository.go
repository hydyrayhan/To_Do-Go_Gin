package repository

import (
	"context"
	"time"
	"todo_api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUser(pool *pgxpool.Pool, user *models.User) (*models.User, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	var query string = `
		INSERT INTO users (email,password)
		VALUES ($1, $2)
		RETURNING id, email,created_at,updated_at;
	`

	var err error = pool.QueryRow(ctx, query, user.Email, user.Password).Scan(
		&user.Id,
		&user.Email,
		&user.Created_at,
		&user.Updated_at,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByEmail(pool *pgxpool.Pool, email string) (*models.User, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
		SELECT id, email, password, created_at, updated_at 
		FROM users
		WHERE email=$1;
	`

	var user models.User
	err := pool.QueryRow(ctx, query, email).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.Created_at,
		&user.Updated_at,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserById(pool *pgxpool.Pool, id int) (*models.User, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
		SELECT id, email,created_at,updated_at 
		FROM users
		WHERE id=$1;
	`

	var user models.User
	err := pool.QueryRow(ctx, query, id).Scan(
		&user.Id,
		&user.Email,
		&user.Created_at,
		&user.Updated_at,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
