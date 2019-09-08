package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/superhero-suggestions/internal/config"
	"github.com/superhero-suggestions/internal/es"
)

// Controller contains definition of controller methods.
type Controller interface {
	RegisterRoutes() *gin.Engine
}

type controller struct {
	ES *es.ES
}

// NewController returns new controller.
func NewController(cfg *config.Config) (ctrl Controller, err error) {
	e, err := es.NewES(cfg)
	if err != nil {
		return nil, err
	}

	return &controller{
		ES:       e,
	}, nil
}

// RegisterRoutes registers all the superhero_suggestions API routes.
func (ctl *controller) RegisterRoutes() *gin.Engine {
	router := gin.Default()

	sr := router.Group("/api/v1/superhero_suggestions")

	// Middleware.
	// sr.Use(c.Authorize)

	// Routes.
	sr.POST("/suggest", ctl.Suggest)

	return router
}