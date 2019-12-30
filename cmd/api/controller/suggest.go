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

		fmt.Println("HandleESRequest")
		for _, r := range result {
			fmt.Printf("%+v", r)
			fmt.Println()
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
