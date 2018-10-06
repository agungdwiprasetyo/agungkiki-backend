package main

import (
	"fmt"
	"log"

	"github.com/agungdwiprasetyo/agungkiki-backend/config"
	"github.com/agungdwiprasetyo/agungkiki-backend/middleware"
	"github.com/agungdwiprasetyo/agungkiki-backend/src/presenter"
	"github.com/agungdwiprasetyo/agungkiki-backend/src/repository"
	"github.com/agungdwiprasetyo/agungkiki-backend/src/usecase"
	"github.com/labstack/echo"
)

// Service main model
type Service struct {
	conf        config.Config
	httpHandler *presenter.InvitationPresenter
}

// NewService create new service
func NewService(conf config.Config) *Service {
	repositoryDecorator := repository.NewInvitationRepository(conf.LoadDB())

	uc := usecase.NewInvitationUsecase(repositoryDecorator)

	service := new(Service)
	service.conf = conf
	service.httpHandler = presenter.NewInvitationPresenter(uc)
	return service
}

// ServeHTTP service
func (serv *Service) ServeHTTP(port int) {
	app := echo.New()
	app.Use(middleware.SetCORS())

	storeGroup := app.Group("/invitation")
	serv.httpHandler.Mount(storeGroup)

	appPort := fmt.Sprintf(":%d", port)
	if err := app.Start(appPort); err != nil {
		log.Fatal(err)
	}
}
