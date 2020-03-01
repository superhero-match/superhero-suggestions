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

	ctrl "github.com/superhero-suggestions/cmd/api/model"
)

// Suggest returns the suggestions for the Superhero.
func (ctl *Controller) Suggest(c *gin.Context) {
	var req ctrl.Request

	err := c.BindJSON(&req)
	if checkError(err, c) {
		ctl.Service.Logger.Error(
			"failed to bind JSON to value of type Request",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.Service.TimeFormat)),
		)

		return
	}

	if req.IsESRequest {
		result, esSuperheroIDs, err := ctl.Service.HandleESRequest(req)
		if checkError(err, c) {
			ctl.Service.Logger.Error(
				"failed while executing service.HandleESRequest()",
				zap.String("err", err.Error()),
				zap.String("time", time.Now().UTC().Format(ctl.Service.TimeFormat)),
			)

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":       http.StatusOK,
			"suggestions":  result,
			"superheroIds": esSuperheroIDs,
		})

		return
	}

	result, err := ctl.Service.GetCachedSuggestions(req)
	if checkError(err, c) {
		ctl.Service.Logger.Error(
			"failed while executing service.GetCachedSuggestions()",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.Service.TimeFormat)),
		)

		return
	}

	fmt.Println("GetCachedSuggestions")
	for _, r := range result {
		fmt.Printf("%+v", r)
		fmt.Println()
	}


	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"suggestions":  result,
		"superheroIds": make([]string, 0),
	})
}

func checkError(err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":       http.StatusInternalServerError,
			"suggestions":  make([]ctrl.Superhero, 0),
			"superheroIds": make([]string, 0),
		})

		return true
	}

	return false
}
