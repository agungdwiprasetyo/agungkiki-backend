package main

import (
	"fmt"
	"log"
	"time"

	"github.com/agungdwiprasetyo/agungkiki-backend/config"
	"github.com/agungdwiprasetyo/agungkiki-backend/middleware"
	"github.com/agungdwiprasetyo/agungkiki-backend/src/presenter"
	"github.com/agungdwiprasetyo/agungkiki-backend/src/repository"
	"github.com/agungdwiprasetyo/agungkiki-backend/src/usecase"
	tokenModule "github.com/agungdwiprasetyo/agungkiki-backend/token"
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
	token := tokenModule.New(conf.LoadPrivateKey(), conf.LoadPublicKey(), 12*time.Hour)

	uc := usecase.NewInvitationUsecase(token, repositoryDecorator)

	service := new(Service)
	service.conf = conf
	service.httpHandler = presenter.NewInvitationPresenter(uc)
	return service
}

// ServeHTTP service
func (serv *Service) ServeHTTP(port int) {
	app := echo.New()
	app.Use(middleware.SetCORS(), middleware.Logger())

	storeGroup := app.Group("/invitation")
	serv.httpHandler.Mount(storeGroup)

	appPort := fmt.Sprintf(":%d", port)
	if err := app.Start(appPort); err != nil {
		log.Fatal(err)
	}
}
