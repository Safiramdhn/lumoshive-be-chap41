package infra

import (
	"lumoshive-be-chap41/config"
	"lumoshive-be-chap41/controller"
	"lumoshive-be-chap41/database"
	"lumoshive-be-chap41/log"
	"lumoshive-be-chap41/middleware"
	"lumoshive-be-chap41/repository"
	"lumoshive-be-chap41/service"

	"go.uber.org/zap"
)

type ServiceContext struct {
	Cfg        config.Config
	Ctl        controller.Controller
	Log        *zap.Logger
	Cacher     database.Cacher
	Middleware middleware.Middleware
}

func NewServiceContext() (*ServiceContext, error) {

	handlerError := func(err error) (*ServiceContext, error) {
		return nil, err
	}

	// instance config
	config, err := config.LoadConfig()
	if err != nil {
		handlerError(err)
	}

	// instance looger
	log, err := log.InitZapLogger(config)
	if err != nil {
		handlerError(err)
	}

	// instance database
	db, err := database.ConnectDB(config)
	if err != nil {
		handlerError(err)
	}

	rdb := database.NewCacher(config, 60*60)

	// instance repository
	repository := repository.NewRepository(db)

	// instance service
	service := service.NewService(repository)

	// instance controller
	Ctl := controller.NewController(service, log, rdb)

	middleware := middleware.NewMiddleware(rdb)

	return &ServiceContext{Cfg: config, Ctl: *Ctl, Log: log, Cacher: rdb, Middleware: middleware}, nil
}
