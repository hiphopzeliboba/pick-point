package pickpoint

import (
	"context"
	"pickpoint/internal/model"
	"pickpoint/internal/repository"
)

type PickPointService struct {
	repo repository.PickPointRepository
}

func NewPickPointService(pickPointRepo repository.PickPointRepository) *PickPointService {
	return &PickPointService{
		repo: pickPointRepo,
	}
}

func (s *PickPointService) CreatePickPoint(ctx context.Context, pp *model.PickPoint) (*model.PickPoint, error) {
	if !model.IsValidCity(pp.City) {
		return nil, model.ErrInvalidCity
	}
	return s.repo.CreatePickPoint(ctx, pp)
}

func (s *PickPointService) GetById(ctx context.Context, id int) (*model.PickPoint, error) {
	return s.repo.GetPickPointById(ctx, id)
}

func (s *PickPointService) List(ctx context.Context, limit, offset int) ([]*model.PickPoint, error) {
	return s.repo.List(ctx, limit, offset)
}
