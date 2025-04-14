package user

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"pickpoint/internal/client/db"
	"pickpoint/internal/model"
)

const (
	tableName = "users"

	idColumn        = "id"
	emailColumn     = "email"
	passwordColumn  = "password"
	roleColumn      = "role"
	createdAtColumn = "created_at"
)

type repo struct {
	db db.Client
}

func NewUserRepository(db db.Client) *repo {
	return &repo{db: db}
}

func (r *repo) CreateUser(ctx context.Context, u *model.User) (*model.User, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(emailColumn, passwordColumn, roleColumn).
		Values(u.Email, u.Password, u.Role).
		Suffix("RETURNING id, email, role, created_at")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var result model.User
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&result.ID,
		&result.Email,
		&result.Role,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	builder := sq.Select(idColumn, emailColumn, passwordColumn, roleColumn, createdAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{emailColumn: email})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.GetUserByEmail",
		QueryRaw: query,
	}

	var result model.User
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&result.ID,
		&result.Email,
		&result.Password,
		&result.Role,
		&result.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, model.ErrNotFound
		}
		return nil, err
	}

	return &result, nil
}
