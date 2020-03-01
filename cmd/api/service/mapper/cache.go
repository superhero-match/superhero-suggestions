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
package mapper

import (
	"github.com/superhero-suggestions/cmd/api/model"
	cache "github.com/superhero-suggestions/internal/cache/model"
)

// MapCacheSuggestionsToResult maps cache Superhero models to API models that are returned to the user.
func MapCacheSuggestionsToResult(cachedSuggestions []cache.Superhero) (result []model.Superhero) {
	for _, s := range cachedSuggestions {
		superhero := model.Superhero{
			ID:                s.ID,
			SuperheroName:     s.SuperheroName,
			MainProfilePicURL: s.MainProfilePicURL,
			Gender:            s.Gender,
			Age:               s.Age,
			Lat:               s.Lat,
			Lon:               s.Lon,
			Birthday:          s.Birthday,
			Country:           s.Country,
			City:              s.City,
			SuperPower:        s.SuperPower,
			AccountType:       s.AccountType,
			CreatedAt:         s.CreatedAt,
		}

		for _, profilePicture := range s.ProfilePictures {
			superhero.ProfilePictures = append(superhero.ProfilePictures, model.ProfilePicture{
				ID:                profilePicture.ID,
				SuperheroID:       profilePicture.SuperheroID,
				ProfilePictureURL: profilePicture.ProfilePictureURL,
				Position:          profilePicture.Position,
			})
		}

		result = append(result, superhero)
	}

	return result
}

// MapCacheSuggestionToResult maps cache Superhero model to API model that is returned to the user.
func MapCacheSuggestionToResult(s cache.Superhero) model.Superhero {
	superhero := model.Superhero{
		ID:                s.ID,
		SuperheroName:     s.SuperheroName,
		MainProfilePicURL: s.MainProfilePicURL,
		Gender:            s.Gender,
		Age:               s.Age,
		Lat:               s.Lat,
		Lon:               s.Lon,
		Birthday:          s.Birthday,
		Country:           s.Country,
		City:              s.City,
		SuperPower:        s.SuperPower,
		AccountType:       s.AccountType,
		CreatedAt:         s.CreatedAt,
	}

	for _, profilePicture := range s.ProfilePictures {
		superhero.ProfilePictures = append(superhero.ProfilePictures, model.ProfilePicture{
			ID:                profilePicture.ID,
			SuperheroID:       profilePicture.SuperheroID,
			ProfilePictureURL: profilePicture.ProfilePictureURL,
			Position:          profilePicture.Position,
		})
	}

	return superhero
}

// MapResultToCacheSuggestions maps API Superhero models to Cache Superhero models.
func MapResultToCacheSuggestions(result []model.Superhero) (cacheSuggestions []cache.Superhero) {
	for _, s := range result {
		cacheSuggestion := cache.Superhero{
			ID:                s.ID,
			SuperheroName:     s.SuperheroName,
			MainProfilePicURL: s.MainProfilePicURL,
			Gender:            s.Gender,
			Age:               s.Age,
			Lat:               s.Lat,
			Lon:               s.Lon,
			Birthday:          s.Birthday,
			Country:           s.Country,
			City:              s.City,
			SuperPower:        s.SuperPower,
			AccountType:       s.AccountType,
			CreatedAt:         s.CreatedAt,
		}

		for _, profilePicture := range s.ProfilePictures {
			cacheSuggestion.ProfilePictures = append(cacheSuggestion.ProfilePictures, cache.ProfilePicture{
				ID:                profilePicture.ID,
				SuperheroID:       profilePicture.SuperheroID,
				ProfilePictureURL: profilePicture.ProfilePictureURL,
				Position:          profilePicture.Position,
			})
		}

		cacheSuggestions = append(cacheSuggestions, cacheSuggestion)
	}

	return cacheSuggestions
}
