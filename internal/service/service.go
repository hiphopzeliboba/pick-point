package service

import (
	"context"
	"pickpoint/internal/model"
)

type IntakeService interface {
	CreateIntake(ctx context.Context, intake *model.Intake) (*model.Intake, error)
	CloseIntake(ctx context.Context, intakeId int) error
}

type PickPointService interface {
	CreatePickPoint(ctx context.Context, pickPoint *model.PickPoint) (*model.PickPoint, error)
	GetById(ctx context.Context, id int) (*model.PickPoint, error)
	List(ctx context.Context, limit, offset int) ([]*model.PickPoint, error)
}

type UserService interface {
	CreateUser(ctx context.Context, email, password string, role model.Role) (*model.User, error)
	Login(ctx context.Context, email, password string) (*model.User, error)
}
