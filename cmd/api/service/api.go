package service

import (
	"fmt"
	ctrl "github.com/superhero-suggestions/cmd/api/model"
	"github.com/superhero-suggestions/cmd/api/service/mapper"
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
	fmt.Println("CacheSuggestions err: ")
	fmt.Println(err)

	if err != nil {
		return nil, nil, err
	}

	// Return max 10 suggestions.
	suggestions = mapper.CutTotalResultToPageSize(srv.PageSize, result)

	// Remove first 10 Superhero ids as these suggestions already being returned with the first batch.
	esSuperheroIDs = mapper.CutFirstPageIdsFromESSuperheroIDs(srv.PageSize,  esSuperheroIDs)

	if suggestions == nil {
		suggestions = make([]ctrl.Superhero, 0)
	}

	if esSuperheroIDs == nil {
		esSuperheroIDs = make([]string, 0)
	}

	if len(suggestions) > 0 {
		keys := make([]string, 0)

		for _, res := range suggestions {
			keys = append(keys, fmt.Sprintf("choice.%s.%s", res.ID, req.ID))
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

	return suggestions, esSuperheroIDs, nil
}
