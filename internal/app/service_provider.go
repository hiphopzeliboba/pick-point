package app

import (
	"context"
	"log"
	"pickpoint/internal/api/handler"
	"pickpoint/internal/client/db"
	"pickpoint/internal/client/db/pg"
	"pickpoint/internal/closer"
	"pickpoint/internal/config"
	"pickpoint/internal/repository"
	"pickpoint/internal/repository/intake"
	"pickpoint/internal/repository/pickpoint"
	"pickpoint/internal/repository/user"
	"pickpoint/internal/service"
	intakeService "pickpoint/internal/service/intake"
	pickpointService "pickpoint/internal/service/pickpoint"
	userService "pickpoint/internal/service/user"
)

type serviceProvider struct {
	pgConfig config.PGConfig

	dbClient db.Client

	pickpointRepository repository.PickPointRepository
	intakeRepository    repository.IntakeRepository
	userRepository      repository.UserRepository

	pickpointService service.PickPointService
	intakeService    service.IntakeService
	userService      service.UserService

	pickpointHandler handler.PickPointHandler
	intakeHandler    handler.IntakeHandler
	userHandler      handler.UserHandler
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().CONN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) PickPointRepository(ctx context.Context) repository.PickPointRepository {
	if s.pickpointRepository == nil {
		s.pickpointRepository = pickpoint.NewPickPointRepository(s.DBClient(ctx))
	}
	return s.pickpointRepository
}

func (s *serviceProvider) IntakeRepository(ctx context.Context) repository.IntakeRepository {
	if s.intakeRepository == nil {
		s.intakeRepository = intake.NewIntakeRepository(s.DBClient(ctx))
	}
	return s.intakeRepository
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = user.NewUserRepository(s.DBClient(ctx))
	}
	return s.userRepository
}

func (s *serviceProvider) PickPointService(ctx context.Context) service.PickPointService {
	if s.pickpointService == nil {
		s.pickpointService = pickpointService.NewPickPointService(s.pickpointRepository)
	}
	return s.pickpointService
}
func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewUserService(s.userRepository)
	}
	return s.userService
}
func (s *serviceProvider) IntakeService(ctx context.Context) service.IntakeService {
	if s.intakeService == nil {
		s.intakeService = intakeService.NewIntakeService(s.intakeRepository)
	}
	return s.intakeService
}
