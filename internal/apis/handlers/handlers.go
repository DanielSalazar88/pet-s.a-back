package handlers

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"pet-s.a/internal/service"
)

type Handlers interface {
	StartServer(e *echo.Echo, address string) error
	registerRoutes(e *echo.Echo)

	insertClient(c echo.Context) error
}

type MyHandlers struct {
	DataValidator validator.Validate
	Service       service.Service
}

func New(s service.Service) Handlers {
	return &MyHandlers{
		DataValidator: *validator.New(),
		Service:       s,
	}
}

func (h *MyHandlers) StartServer(e *echo.Echo, address string) error {

	errCh := make(chan error, 1)

	go func() {
		h.registerRoutes(e)

		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{echo.GET, echo.POST},
			AllowHeaders:     []string{echo.HeaderContentType, echo.HeaderAuthorization},
			AllowCredentials: true,
		}))

		err := e.Start(address)
		errCh <- err
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errCh:
		if err != nil {
			log.Printf("Error al iniciar el servidor: %s\n", err)
		} else {
			log.Println("Servidor detenido sin error.")
		}

	case <-stop:
		log.Println("Señal de interrupción recibida. Deteniendo el servidor.")
		e.Shutdown(context.Background())
	}

	return nil
}
