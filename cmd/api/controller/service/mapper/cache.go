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
