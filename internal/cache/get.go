package cache

import (
	"github.com/superhero-suggestions/internal/cache/model"
	"reflect"
)

// GetSuggestions fetches suggestions from cache.
func (c *Cache) GetSuggestions(keys string) (suggestions []model.Superhero, err error) {
	sliceCmd := c.Redis.MGet(keys)
	if sliceCmd.Err() != nil {
		return nil, err
	}

	result, err := sliceCmd.Result()
	if err != nil {
		return nil, err
	}

	for _, v := range result {
		item := reflect.ValueOf(v)
		var suggestion model.Superhero

		suggestion.ID = item.Field(0).Interface().(string)
		suggestion.SuperheroName = item.Field(1).Interface().(string)
		suggestion.MainProfilePicURL = item.Field(2).Interface().(string)
		suggestion.ProfilePictures = item.Field(3).Interface().([]model.ProfilePicture)
		suggestion.Gender = item.Field(4).Interface().(int)
		suggestion.Age = item.Field(5).Interface().(int)
		suggestion.Lat = item.Field(6).Interface().(float64)
		suggestion.Lon = item.Field(7).Interface().(float64)
		suggestion.Birthday = item.Field(8).Interface().(string)
		suggestion.Country = item.Field(9).Interface().(string)
		suggestion.City = item.Field(10).Interface().(string)
		suggestion.SuperPower = item.Field(11).Interface().(string)
		suggestion.AccountType = item.Field(12).Interface().(string)
		suggestion.CreatedAt = item.Field(13).Interface().(string)

		suggestions = append(suggestions, suggestion)
	}

	return suggestions, nil
}
