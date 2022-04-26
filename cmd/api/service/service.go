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
	ctrl "github.com/superhero-match/superhero-suggestions/cmd/api/model"
	"github.com/superhero-match/superhero-suggestions/internal/cache"
	"github.com/superhero-match/superhero-suggestions/internal/es"
	"github.com/superhero-match/superhero-suggestions/internal/es/model"
)

// Service interface defines service methods.
type Service interface {
	HandleESRequest(req ctrl.Request, likeSuperheroIDs []string) (suggestions []ctrl.Superhero, esSuperheroIDs []string, err error)
	GetCachedSuggestions(req ctrl.Request) (result []ctrl.Superhero, err error)
	CacheSuggestions(result []ctrl.Superhero) error
	GetCachedChoices(keys []string) (map[string]bool, error)
	GetCachedChoice(key string) (map[string]bool, error)
	GetLikes(superheroID string) ([]string, error)
	DeleteLikes(superheroID string) error
	GetESSuggestions(req ctrl.Request, likeSuperheroIDs []string) (superheros []model.Superhero, err error)
}

// service holds all the different services that are used when handling request.
type service struct {
	ES              es.ES
	Cache           cache.Cache
	PageSize        int
	ChoiceKeyFormat string
}

// New creates value of type Service.
func New(e es.ES, c cache.Cache, pageSize int, choiceKeyFormat string) Service {
	return &service{
		ES:              e,
		Cache:           c,
		PageSize:        pageSize,
		ChoiceKeyFormat: choiceKeyFormat,
	}
}
