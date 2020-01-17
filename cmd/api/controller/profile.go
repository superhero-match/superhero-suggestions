package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	ctrl "github.com/superhero-suggestions/cmd/api/model"
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
