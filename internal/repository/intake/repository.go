package intake

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"pickpoint/internal/client/db"
	"pickpoint/internal/model"
	"pickpoint/internal/repository"
)

const (
	tableName = "intakes"

	idColumn          = "id"
	pickPointIdColumn = "pick_point_id"
	employeeIdColumn  = "employee_id"
	createdAtColumn   = "created_at"
	statusColumn      = "status"
)

type repo struct {
	db db.Client
}

func NewIntakeRepository(db db.Client) repository.IntakeRepository {
	return &repo{db: db}
}

func (r *repo) CreateIntake(ctx context.Context, intake *model.Intake) (*model.Intake, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(pickPointIdColumn, employeeIdColumn, statusColumn).
		Values(intake.PickPointId, intake.EmployeeId, intake.Status).
		Suffix("RETURNING id, pick_up_point_id, employee_id, created_at")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "intake_repository.Create",
		QueryRaw: query,
	}

	var rec model.Intake
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&rec.ID,
		&rec.PickPointId,
		&rec.EmployeeId,
		&rec.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &rec, nil
}

func (r *repo) GetIntakeById(ctx context.Context, id int) (*model.Intake, error) {
	builder := sq.Select(idColumn, pickPointIdColumn, employeeIdColumn, createdAtColumn, statusColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "intake_repository.Get",
		QueryRaw: query,
	}

	var rec model.Intake
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&rec.ID,
		&rec.PickPointId,
		&rec.EmployeeId,
		&rec.CreatedAt,
		&rec.Status,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return &rec, nil
}

func (r *repo) UpdateIntakeStatus(ctx context.Context, intakeId int, status string) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(statusColumn, status).
		Where(sq.Eq{idColumn: intakeId})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "intake_repository.UpdateStatus",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) GetActiveIntakeByPickPoint(ctx context.Context, pickPointId int) (*model.Intake, error) {
	builder := sq.Select(idColumn, pickPointIdColumn, employeeIdColumn, createdAtColumn, statusColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{
			pickPointIdColumn: pickPointId,
			statusColumn:      "in_progress",
		}).
		OrderBy(createdAtColumn + " DESC").
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "intake_repository.GetActiveIntakeByPickPoint",
		QueryRaw: query,
	}

	var rec model.Intake
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&rec.ID,
		&rec.PickPointId,
		&rec.EmployeeId,
		&rec.CreatedAt,
		&rec.Status,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return &rec, nil
}
