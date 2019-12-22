package controller

import (
	"fmt"
	"net/http"

	ctrl "github.com/superhero-suggestions/cmd/api/model"
	"github.com/superhero-suggestions/internal/es/model"

	"github.com/gin-gonic/gin"
)

// Suggest returns the suggestions for the Superhero.
func (ctl *controller) Suggest(c *gin.Context) {
	var req ctrl.Request
	var result []ctrl.Superhero

	//body := c.Request.Body
	//x, _ := ioutil.ReadAll(body)
	//
	//fmt.Printf("%s \n", string(x))


	err := c.BindJSON(&req)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"suggestions": []ctrl.Superhero{},
		})

		return
	}

	superheros, err := ctl.ES.GetSuggestions(
		&model.Request{
			ID:               req.ID,
			LookingForGender: req.LookingForGender,
			Gender:           req.Gender,
			LookingForAgeMin: req.LookingForAgeMin,
			LookingForAgeMax: req.LookingForAgeMax,
			MaxDistance:      req.MaxDistance,
			DistanceUnit:     req.DistanceUnit,
			Lat:              req.Lat,
			Lon:              req.Lon,
			Offset:           req.Offset,
			Size:             req.Size,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"suggestions": []ctrl.Superhero{},
		})

		return
	}

	for _, s := range superheros {
		result = append(result, ctrl.Superhero{
			ID:                s.ID,
			SuperheroName:     s.SuperheroName,
			MainProfilePicURL: s.MainProfilePicURL,
			Gender:            s.Gender,
			Age:               s.Age,
			Lat:               s.Location.Lat,
			Lon:               s.Location.Lon,
			Birthday:          s.Birthday,
			Country:           s.Country,
			City:              s.City,
			SuperPower:        s.SuperPower,
			AccountType:       s.AccountType,
			CreatedAt:         s.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"suggestions": result,
	})
}
