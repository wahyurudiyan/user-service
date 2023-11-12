package users

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/wahyurudiyan/super-sharing/user-svc/internal/core/entity"
	repoIface "github.com/wahyurudiyan/super-sharing/user-svc/internal/core/ports/repository"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repoIface.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) error {
	query := `
	INSERT INTO user (
		unique_id, firstname, lastname, email, password, created_at, updated_at
	)
	VALUES (
		:unique_id, :firstname, :lastname, :email, :password, :created_at, :updated_at
	)`
	_, err := r.db.NamedExecContext(ctx, query, &user)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindUsers(ctx context.Context, limit uint) (entity.Users, error) {
	query := `
	SELECT
		unique_id, firstname, lastname, email, password, created_at, updated_at, deleted_at
	FROM 
		user 
	WHERE 
		deleted_at IS NULL
	LIMIT ?`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var users entity.Users
	err = stmt.Select(&users, limit)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) FindUserByEmail(ctx context.Context, email string) (entity.User, error) {
	query := `
	SELECT
		unique_id, firstname, lastname, email, password, created_at, updated_at
	FROM 
		user 
	WHERE 
		email = ? AND deleted_at IS NULL`
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return entity.User{}, err
	}

	var user entity.User
	err = stmt.Select(&user, email)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) FindUserByUniqueID(ctx context.Context, uniqueId string) (entity.Users, error) {
	query := `SELECT 
		unique_id, firstname, lastname, email, password, created_at, updated_at
		FROM user WHERE unique_id = :unique_id AND deleted_at IS NULL`
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}

	args := entity.User{
		UniqueId: uniqueId,
	}

	var users entity.Users
	err = stmt.Select(&users, args)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) error {
	query := `UPDATE user SET 
	firstname = :firstname,
	lastname = :lastname,
	email = :email,
	password = :password WHERE unique_id = :unique_id AND deleted_at IS NULL`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, user entity.User) error {
	query := `UPDATE user SET deleted_at = :deleted_at WHERE unique_id = :unique_id`
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return err
	}

	user.DeletedAt = time.Now()
	_, err = stmt.ExecContext(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
