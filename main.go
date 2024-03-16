package main

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"pet-s.a/internal/apis/handlers"
	"pet-s.a/internal/database"
	"pet-s.a/internal/repository"
	"pet-s.a/internal/service"
	"pet-s.a/settings"
)

func main() {
	app := fx.New(
		fx.Provide(
			settings.New,
			database.New,
			repository.New,
			handlers.New,
			echo.New,
			service.New,
		),

		fx.Invoke(
			setLifeCycle,
		),
	)

	app.Run()
}

func setLifeCycle(lc fx.Lifecycle, e *echo.Echo, h handlers.Handlers, s *settings.MySettings) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%s", s.Handlers.Port)
			go h.StartServer(e, address)
			return nil
		},

		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
