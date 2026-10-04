package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kmaxsoul/tabaa-store-api/models"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (first_name, last_name, email, password_hash, phone_number)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, role, created_at, updated_at
	`

	err := r.db.QueryRow(
		ctx,
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PasswordHash,
		user.PhoneNumber,
	).Scan(&user.ID, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	return err

}
