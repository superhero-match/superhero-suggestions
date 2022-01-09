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
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"

	ctrl "github.com/superhero-match/superhero-suggestions/cmd/api/model"
	"github.com/superhero-match/superhero-suggestions/internal/cache"
	cm "github.com/superhero-match/superhero-suggestions/internal/cache/model"
	"github.com/superhero-match/superhero-suggestions/internal/config"
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
	FetchAuth(authD *cm.AccessDetails) (string, error)
	GetESSuggestions(req ctrl.Request, likeSuperheroIDs []string) (superheros []model.Superhero, err error)
	ExtractToken(r *http.Request) string
	VerifyToken(r *http.Request) (*jwt.Token, error)
	ExtractTokenMetadata(r *http.Request) (*cm.AccessDetails, error)
}

// service holds all the different services that are used when handling request.
type service struct {
	ES              es.ES
	Cache           cache.Cache
	PageSize        int
	AccessSecret    string
	ChoiceKeyFormat string
}

// NewService creates value of type Service.
func NewService(cfg *config.Config) (Service, error) {
	e, err := es.NewES(cfg)
	if err != nil {
		return nil, err
	}

	c, err := cache.NewCache(cfg)
	if err != nil {
		return nil, err
	}

	return &service{
		ES:              e,
		Cache:           c,
		PageSize:        cfg.App.PageSize,
		AccessSecret:    cfg.JWT.AccessTokenSecret,
		ChoiceKeyFormat: cfg.Cache.ChoiceKeyFormat,
	}, nil
}
