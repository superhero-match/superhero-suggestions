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
package service

import (
	"fmt"
	"github.com/superhero-match/superhero-suggestions/cmd/api/model"
	"github.com/superhero-match/superhero-suggestions/cmd/api/service/mapper"
	"sort"
)

// GetCachedSuggestions fetches suggestions from cache and maps them into result.
func (srv *Service) GetCachedSuggestions(req model.Request) (result []model.Superhero, err error) {
	cachedSuggestions, err := srv.Cache.GetSuggestions(req.SuperheroIDs)
	if err != nil {
		return nil, err
	}

	result = mapper.MapCacheSuggestionsToResult(cachedSuggestions)

	keys := make([]string, 0)

	for _, res := range result {
		keys = append(keys, fmt.Sprintf(srv.Cache.ChoiceKeyFormat, res.ID, req.ID))
	}

	choices, err := srv.GetCachedChoices(keys)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(result); i++ {
		_, ok := choices[result[i].ID]
		if !ok {
			result[i].HasLikedMe = false

			continue
		}

		result[i].HasLikedMe = true
	}

	for l := 0; l < len(result); l++ {
		if len(result[l].ProfilePictures) > 0 {
			sort.Slice(result[l].ProfilePictures, func(i, j int) bool {
				return result[l].ProfilePictures[i].Position < result[l].ProfilePictures[j].Position
			})
		}
	}

	return result, nil
}

// CacheSuggestions maps ES models to Cache models and caches the suggestions.
func (srv *Service) CacheSuggestions(result []model.Superhero) error {
	return srv.Cache.SetSuggestions(mapper.MapResultToCacheSuggestions(result))
}

// GetCachedChoices fetches choices from cache.
func (srv *Service) GetCachedChoices(keys []string) (map[string]bool, error) {
	cachedChoices, err := srv.Cache.GetChoices(keys)
	if err != nil {
		return nil, err
	}

	result := make(map[string]bool)

	for _, cachedChoice := range cachedChoices {
		result[cachedChoice.SuperheroID] = true
	}

	return result, nil
}

// GetCachedChoice fetches choice from cache.
func (srv *Service) GetCachedChoice(key string) (map[string]bool, error) {
	cachedChoice, err := srv.Cache.GetChoice(key)
	if err != nil {
		return nil, err
	}

	result := make(map[string]bool)

	result[cachedChoice.SuperheroID] = true

	return result, nil
}

// GetLikes fetches all the user ids of users who liked this user.
func (srv *Service) GetLikes(superheroID string) ([]string, error) {
	return srv.Cache.GetLikes(superheroID)
}

// DeleteLikes deletes all the user ids of users who liked this user after the results were fetched and processed.
func (srv *Service) DeleteLikes(superheroID string) error {
	// Delete the likes as they were already included in the Elasticsearch query.
	// No need to be fetching the same users over and over again.
	return srv.Cache.DeleteLikes(superheroID)
}