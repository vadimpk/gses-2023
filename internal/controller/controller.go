package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vadimpk/gses-2023/config"
	"github.com/vadimpk/gses-2023/internal/service"
	"github.com/vadimpk/gses-2023/pkg/logging"
	"net/http"
)

type Options struct {
	Services service.Services
	Config   *config.Config
	Logger   logging.Logger
}

type routerContext struct {
	services service.Services
	cfg      *config.Config
	logger   logging.Logger
}

type routerOptions struct {
	router   *gin.Engine
	services service.Services
	cfg      *config.Config
	logger   logging.Logger
}

func New(opts *Options) http.Handler {
	r := gin.Default()

	routerOptions := routerOptions{
		router:   r,
		services: opts.Services,
		cfg:      opts.Config,
		logger:   opts.Logger.Named("HTTPController"),
	}

	setupEmailRoutes(&routerOptions)
	setupCryptoRoutes(&routerOptions)

	return r
}
