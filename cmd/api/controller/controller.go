/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/superhero-match/superhero-suggestions/cmd/api/service"
)

// Controller holds the Controller data.
type Controller struct {
	Service    service.Service
	Logger     *zap.Logger
	TimeFormat string
}

// New returns new controller.
func New(s service.Service, logger *zap.Logger, timeFormat string) *Controller {
	return &Controller{
		Service:    s,
		Logger:     logger,
		TimeFormat: timeFormat,
	}
}

// RegisterRoutes registers all the superhero_suggestions API routes.
func (ctl *Controller) RegisterRoutes() *gin.Engine {
	router := gin.Default()

	sr := router.Group("/api/v1/superhero_suggestions")

	// Routes.
	sr.POST("/suggest", ctl.Suggest)
	sr.GET("/health", ctl.Health)

	return router
}
