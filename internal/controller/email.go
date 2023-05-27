package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type emailRoutes struct {
	routerContext
}

func setupEmailRoutes(opts *routerOptions) {
	emailRoutes := emailRoutes{
		routerContext: routerContext{
			services: opts.services,
			cfg:      opts.cfg,
			logger:   opts.logger.Named("Email"),
		},
	}

	opts.router.POST("/subscribe", emailRoutes.subscribe)
	opts.router.POST("/sendEmails", emailRoutes.sendRateInfo)
}

func (r *emailRoutes) subscribe(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (r *emailRoutes) sendRateInfo(c *gin.Context) {
	c.Status(http.StatusOK)
}
