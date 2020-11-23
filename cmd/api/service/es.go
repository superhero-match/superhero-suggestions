/*
  Copyright (C) 2019 - 2021 MWSOFT
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
package service

import (
	ctrl "github.com/superhero-match/superhero-suggestions/cmd/api/model"
	"github.com/superhero-match/superhero-suggestions/internal/es/model"
)

// GetESSuggestions fetches suggestions from Elasticsearch.
func (srv *Service) GetESSuggestions(req ctrl.Request, likeSuperheroIDs []string) (superheros []model.Superhero, err error) {
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
		likeSuperheroIDs,
	)
	if err != nil {
		return nil, err
	}

	return superheros, nil
}
