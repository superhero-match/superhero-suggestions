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

// GetESSuggestion fetches suggestions from Elasticsearch.
func (srv *Service) GetESSuggestion(superheroID string) (*ctrl.Superhero, error) {
	s, err := srv.ES.GetSuggestion(superheroID)
	if err != nil {
		return nil, err
	}

	superhero := ctrl.Superhero{
		ID:                s.ID,
		SuperheroName:     s.SuperheroName,
		MainProfilePicURL: s.MainProfilePicURL,
		ProfilePictures:   make([]ctrl.ProfilePicture, 0),
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
	}

	for _, profilePicture := range s.ProfilePictures {
		superhero.ProfilePictures = append(superhero.ProfilePictures, ctrl.ProfilePicture{
			ID:                profilePicture.ID,
			SuperheroID:       profilePicture.SuperheroID,
			ProfilePictureURL: profilePicture.ProfilePictureURL,
			Position:          profilePicture.Position,
		})
	}

	return &superhero, nil
}
