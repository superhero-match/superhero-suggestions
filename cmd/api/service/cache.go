package service

import (
	"fmt"
	"github.com/superhero-suggestions/cmd/api/model"
	"github.com/superhero-suggestions/cmd/api/service/mapper"
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
		keys = append(keys, fmt.Sprintf("choice.%s.%s", res.ID, req.ID))
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

	return result, nil
}

// GetCachedSuggestion fetches suggestion from cache and maps it into result.
func (srv *Service) GetCachedSuggestion(key string) (*model.Superhero, error) {
	cachedSuggestion, err := srv.Cache.GetSuggestion(key)
	if err != nil {
		return nil, err
	}

	result := mapper.MapCacheSuggestionToResult(*cachedSuggestion)

	choice, err := srv.GetCachedChoice(fmt.Sprintf("choice.%s.%s", result.ID, key))
	if err != nil {
		return nil, err
	}

	_, ok := choice[result.ID]
	if !ok {
		result.HasLikedMe = false

		return &result, nil
	}

	result.HasLikedMe = true

	return &result, nil
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
