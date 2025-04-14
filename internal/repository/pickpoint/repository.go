package pickpoint

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"pickpoint/internal/client/db"
	"pickpoint/internal/model"
	"pickpoint/internal/repository"
)

const (
	tableName = "pick_points"

	idColumn        = "id"
	cityColumn      = "city"
	createdAtColumn = "create_at"
)

type repo struct {
	db db.Client
}

func NewPickPointRepository(db db.Client) repository.PickPointRepository {
	return &repo{db: db}
}

func (r *repo) CreatePickPoint(ctx context.Context, pp *model.PickPoint) (*model.PickPoint, error) {
	builder := sq.
		Insert(tableName).
		Columns(cityColumn).
		Values(pp.City).
		Suffix("RETURNING id, registration_date, city").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "pvz_repository.CreatePVZ",
		QueryRaw: query,
	}

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&pp.ID, &pp.CreatedAt, &pp.City)
	if err != nil {
		return nil, err
	}
	return pp, err

}

func (r *repo) GetPickPointById(ctx context.Context, id int) (*model.PickPoint, error) {
	builder := sq.Select(idColumn, cityColumn, createdAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "pickpoint_repository.Get",
		QueryRaw: query,
	}

	var result model.PickPoint
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&result.ID,
		&result.City,
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

func (r *repo) List(ctx context.Context, limit, offset int) ([]*model.PickPoint, error) {
	builder := sq.Select(idColumn, cityColumn, createdAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		OrderBy(createdAtColumn + " DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset))

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "pickpoint_repository.List",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*model.PickPoint
	for rows.Next() {
		var pp model.PickPoint
		err = rows.Scan(
			&pp.ID,
			&pp.City,
			&pp.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, &pp)
	}

	return results, nil
}
