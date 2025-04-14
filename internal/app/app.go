package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"pickpoint/internal/api/handler/intake"
	"pickpoint/internal/api/handler/pickpoint"
	"pickpoint/internal/api/handler/user"
	r "pickpoint/internal/api/router"
	"pickpoint/internal/closer"
	"pickpoint/internal/config"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runHTTPServer()
	//return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,

		//a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) runHTTPServer() error {
	log.Printf("HTTP server is running on %s", a.httpServer.Addr)

	err := a.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("Failed to start HTTP server: %s", err)
	}

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	userHandler := user.NewUserHandler(a.serviceProvider.userService)
	pickPointHandler := pickpoint.NewPickPointHandler(a.serviceProvider.pickpointService)
	intakeHandler := intake.NewIntakeHandler(a.serviceProvider.intakeService)

	router := r.NewRouter(userHandler, pickPointHandler, intakeHandler)

	a.httpServer = &http.Server{
		Addr:    ":8080", // В реальном приложении нужно брать из конфига
		Handler: router,
	}

	return nil
}
