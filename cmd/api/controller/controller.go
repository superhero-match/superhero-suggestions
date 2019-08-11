package controller

import (
	"github.com/gin-gonic/gin"
)

// Controller contains definition of controller methods.
type Controller interface {
	RegisterRoutes() *gin.Engine
}

type controller struct {
}

// NewController returns new controller.
func NewController() Controller {
	return &controller{}
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