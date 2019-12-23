package service

import (
	"strings"

	"github.com/superhero-suggestions/cmd/api/controller/service/mapper"
	"github.com/superhero-suggestions/cmd/api/model"
)

// GetCachedSuggestions fetches suggestions from cache and maps them into result.
func (srv *Service) GetCachedSuggestions(req model.Request) (result []model.Superhero, err error) {
	cachedSuggestions, err := srv.Cache.GetSuggestions(strings.Join(req.SuperheroIDs, ","))
	if err != nil {
		return nil, err
	}

	return mapper.MapCacheSuggestionsToResult(cachedSuggestions), nil
}

// CacheSuggestions maps ES models to Cache models and caches the suggestions.
func (srv *Service) CacheSuggestions(result []model.Superhero) error {
	return srv.Cache.SetSuggestions(mapper.MapResultToCacheSuggestions(result))
}
