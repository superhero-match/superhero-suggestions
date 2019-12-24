package service

import (
	ctrl "github.com/superhero-suggestions/cmd/api/model"
	"github.com/superhero-suggestions/internal/es/model"
)

// GetESSuggestions fetches suggestions from Elasticsearch.
func (srv *Service) GetESSuggestions(req ctrl.Request) (superheros []model.Superhero, err error) {
	superheros, err = srv.ES.GetSuggestions(
		&model.Request{
			ID:                    req.ID,
			LookingForGender:      req.LookingForGender,
			Gender:                req.Gender,
			LookingForAgeMin:      req.LookingForAgeMin,
			LookingForAgeMax:      req.LookingForAgeMax,
			MaxDistance:           req.MaxDistance,
			DistanceUnit:          req.DistanceUnit,
			Lat:                   req.Lat,
			Lon:                   req.Lon,
			RetrievedSuperheroIDs: req.RetrievedSuperheroIDs,
		},
	)
	if err != nil {
		return nil, err
	}

	return superheros, nil
}
