package repository

import (
	"context"
	"pickpoint/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, u *model.User) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}

type PickPointRepository interface {
	CreatePickPoint(ctx context.Context, pp *model.PickPoint) (*model.PickPoint, error)
	GetPickPointById(ctx context.Context, id int) (*model.PickPoint, error)
	List(ctx context.Context, limit, offset int) ([]*model.PickPoint, error)
}

type IntakeRepository interface { //AcceptanceRepository
	CreateIntake(ctx context.Context, intake *model.Intake) (*model.Intake, error)
	GetActiveIntakeByPickPoint(ctx context.Context, pickPointId int) (*model.Intake, error)
	GetIntakeById(ctx context.Context, id int) (*model.Intake, error)
	UpdateIntakeStatus(ctx context.Context, intakeId int, status string) error
	//CloseIntake(ctx context.Context, PickPoint int) error
	//GetLastIntake(ctx context.Context, PickPoint int) (*model.Intake, error)
}
