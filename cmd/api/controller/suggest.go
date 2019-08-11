package controller

import (
	"net/http"

	ctrl "github.com/superhero-suggestions/cmd/api/model"

	"github.com/gin-gonic/gin"
)

// Suggest returns the suggestions for the Superhero.
func (ctl *controller) Suggest(c *gin.Context) {
	req := new(ctrl.Request)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"suggestions": []ctrl.Suggestion{},
		})

		return
	}

	// TO-DO: fetch the suggestions from ES.

	c.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"suggestions": []ctrl.Suggestion{},
	})
}
