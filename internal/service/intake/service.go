package intake

import (
	"context"
	"errors"
	"pickpoint/internal/model"
	"pickpoint/internal/repository"
)

type IntakeService struct {
	intakeRepo repository.IntakeRepository
}

func NewIntakeService(intakeRepo repository.IntakeRepository) *IntakeService {
	return &IntakeService{
		intakeRepo: intakeRepo,
	}
}

func (s *IntakeService) CreateIntake(ctx context.Context, intake *model.Intake) (*model.Intake, error) {
	// Проверяем, есть ли активная приёмка
	activeIntake, err := s.intakeRepo.GetActiveIntakeByPickPoint(ctx, intake.PickPointId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, err
	}
	if activeIntake != nil {
		return nil, model.ErrActiveIntakeExists
	}

	// Устанавливаем статус по умолчанию
	intake.Status = "in_progress"

	return s.intakeRepo.CreateIntake(ctx, intake)
}

func (s *IntakeService) CloseIntake(ctx context.Context, intakeId int) error {
	statusClose := "close"

	// Получаем текущую приёмку
	intake, err := s.intakeRepo.GetIntakeById(ctx, intakeId)
	if err != nil {
		return err
	}

	// Проверяем статус
	if intake.Status == statusClose {
		return model.ErrIntakeAlreadyClosed
	}

	// Закрываем приёмку
	return s.intakeRepo.UpdateIntakeStatus(ctx, intakeId, statusClose)
}
