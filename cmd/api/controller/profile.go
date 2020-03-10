/*
  Copyright (C) 2019 - 2020 MWSOFT
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
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	ctrl "github.com/superhero-match/superhero-suggestions/cmd/api/model"
)

// Profile returns Superhero data with specified id.
func (ctl *Controller) Profile(c *gin.Context) {
	var req ctrl.ProfileRequest

	err := c.BindJSON(&req)
	if checkProfileRequestError(err, c) {
		ctl.Service.Logger.Error(
			"failed to bind JSON to value of type ProfileRequest",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.Service.TimeFormat)),
		)

		return
	}

	suggestion, err := ctl.Service.GetCachedSuggestion(
		fmt.Sprintf(ctl.Service.Cache.SuggestionKeyFormat, req.SuperheroID),
	)
	if checkProfileRequestError(err, c) {
		ctl.Service.Logger.Error(
			"failed to bind JSON to value of type ProfileRequest",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.Service.TimeFormat)),
		)

		return
	}

	if suggestion != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"suggestion": suggestion,
		})

		return
	}

	suggestion, err = ctl.Service.GetESSuggestion(req.SuperheroID)
	if checkProfileRequestError(err, c) {
		ctl.Service.Logger.Error(
			"failed to bind JSON to value of type ProfileRequest",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.Service.TimeFormat)),
		)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     http.StatusOK,
		"suggestion": suggestion,
	})
}

func checkProfileRequestError(err error, c *gin.Context) bool {
	if err != nil {
		var suggestion ctrl.Superhero
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":     http.StatusInternalServerError,
			"suggestion": suggestion,
		})

		return true
	}

	return false
}
