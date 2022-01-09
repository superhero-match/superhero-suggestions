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
package service

import (
	"fmt"
	"sort"

	ctrl "github.com/superhero-match/superhero-suggestions/cmd/api/model"
	"github.com/superhero-match/superhero-suggestions/cmd/api/service/mapper"
)

// HandleESRequest fetches suggestions from Elasticsearch,
// then caches them and returns page size of results.
func (srv *service) HandleESRequest(req ctrl.Request, likeSuperheroIDs []string) (suggestions []ctrl.Superhero, esSuperheroIDs []string, err error) {
	superheros, err := srv.GetESSuggestions(req, likeSuperheroIDs)
	if err != nil {
		return nil, nil, err
	}

	result, esSuperheroIDs := mapper.MapESSuggestionsToResult(superheros)

	err = srv.CacheSuggestions(result)
	fmt.Println("CacheSuggestions err: ")
	fmt.Println(err)

	if err != nil {
		return nil, nil, err
	}

	// Return max 10 suggestions.
	suggestions = mapper.CutTotalResultToPageSize(srv.PageSize, result)

	// Remove first 10 Superhero ids as these suggestions already being returned with the first batch.
	esSuperheroIDs = mapper.CutFirstPageIdsFromESSuperheroIDs(srv.PageSize, esSuperheroIDs)

	if suggestions == nil {
		suggestions = make([]ctrl.Superhero, 0)
	}

	if esSuperheroIDs == nil {
		esSuperheroIDs = make([]string, 0)
	}

	if len(suggestions) > 0 {
		keys := make([]string, 0)

		for _, res := range suggestions {
			keys = append(keys, fmt.Sprintf(srv.ChoiceKeyFormat, res.ID, req.ID))
		}

		choices, err := srv.GetCachedChoices(keys)
		fmt.Println("GetCachedChoices err: ")
		fmt.Println(err)

		if err != nil {
			return nil, nil, err
		}

		for i := 0; i < len(suggestions); i++ {
			_, ok := choices[suggestions[i].ID]
			if !ok {
				suggestions[i].HasLikedMe = false

				continue
			}

			suggestions[i].HasLikedMe = true
		}
	}

	for l := 0; l < len(suggestions); l++ {
		if len(suggestions[l].ProfilePictures) > 0 {
			sort.Slice(suggestions[l].ProfilePictures, func(i, j int) bool {
				return suggestions[l].ProfilePictures[i].Position < suggestions[l].ProfilePictures[j].Position
			})
		}
	}

	return suggestions, esSuperheroIDs, nil
}
