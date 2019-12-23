package service

import (
	"github.com/superhero-suggestions/cmd/api/controller/service/mapper"
	ctrl "github.com/superhero-suggestions/cmd/api/model"
)

// HandleESRequest fetches suggestions from Elasticsearch,
// then caches them and returns page size of results.
func (srv *Service) HandleESRequest(req ctrl.Request) (suggestions []ctrl.Superhero, esSuperheroIDs []string, err error) {
	superheros, err := srv.GetESSuggestions(req)
	if err != nil {
		return nil, nil, err
	}

	result, esSuperheroIDs := mapper.MapESSuggestionsToResult(superheros)

	err = srv.CacheSuggestions(result)
	if err != nil {
		return nil, nil, err
	}

	// Return max 10 suggestions.
	suggestions = mapper.CutTotalResultToPageSize(srv.PageSize, result)

	return suggestions, esSuperheroIDs, nil
}
