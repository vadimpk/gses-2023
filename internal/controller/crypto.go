package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type cryptoRoutes struct {
	routerContext
}

func setupCryptoRoutes(opts *routerOptions) {
	cryptoRoutes := cryptoRoutes{
		routerContext: routerContext{
			services: opts.services,
			cfg:      opts.cfg,
			logger:   opts.logger.Named("Crypto"),
		},
	}

	opts.router.GET("/rate", cryptoRoutes.getRate)
}

func (r *cryptoRoutes) getRate(c *gin.Context) {
	c.Status(http.StatusOK)
}
