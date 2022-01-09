/*
  Copyright (C) 2019 - 2022 MWSOFT
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
	"github.com/superhero-match/superhero-suggestions/cmd/api/model"
	es "github.com/superhero-match/superhero-suggestions/internal/es/model"
)

// CutTotalResultToPageSize takes only specified page size, default is 10, or less from total result set.
func CutTotalResultToPageSize(pageSize int, totalResult []model.Superhero) (suggestions []model.Superhero) {
	for _, suggestion := range totalResult {
		if pageSize == int(0) {
			break
		}

		suggestions = append(suggestions, suggestion)

		pageSize--
	}

	return suggestions
}

// CutFirstPageIdsFromESSuperheroIDs takes only specified page size, default is 10, or less from total result set.
func CutFirstPageIdsFromESSuperheroIDs(numIDsToSkip int, esSuperheroIDs []string) (result []string) {
	for _, esSuperheroID := range esSuperheroIDs {
		if len(esSuperheroIDs) < numIDsToSkip {
			return result
		}

		if numIDsToSkip != 0 {
			numIDsToSkip--

			continue
		}

		result = append(result, esSuperheroID)
	}

	return result
}

// MapESSuggestionsToResult maps ES Superhero to result Superhero.
func MapESSuggestionsToResult(superheros []es.Superhero) (result []model.Superhero, esSuperheroIDs []string) {
	for _, s := range superheros {
		esSuperheroIDs = append(esSuperheroIDs, s.ID)

		superhero := model.Superhero{
			ID:                s.ID,
			SuperheroName:     s.SuperheroName,
			MainProfilePicURL: s.MainProfilePicURL,
			ProfilePictures:   make([]model.ProfilePicture, 0),
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
			superhero.ProfilePictures = append(superhero.ProfilePictures, model.ProfilePicture{
				ID:                profilePicture.ID,
				SuperheroID:       profilePicture.SuperheroID,
				ProfilePictureURL: profilePicture.ProfilePictureURL,
				Position:          profilePicture.Position,
			})
		}

		result = append(result, superhero)
	}

	return result, esSuperheroIDs
}
