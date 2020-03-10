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
	"go.uber.org/zap"

	"github.com/superhero-match/superhero-suggestions/internal/cache"
	"github.com/superhero-match/superhero-suggestions/internal/config"
	"github.com/superhero-match/superhero-suggestions/internal/es"
)

// Service holds all the different services that are used when handling request.
type Service struct {
	ES         *es.ES
	Cache      *cache.Cache
	PageSize   int
	Logger     *zap.Logger
	TimeFormat string
}

// NewService creates value of type Service.
func NewService(cfg *config.Config) (*Service, error) {
	e, err := es.NewES(cfg)
	if err != nil {
		return nil, err
	}

	c, err := cache.NewCache(cfg)
	if err != nil {
		return nil, err
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	defer logger.Sync()

	return &Service{
		ES:         e,
		Cache:      c,
		PageSize:   cfg.App.PageSize,
		Logger:     logger,
		TimeFormat: cfg.App.TimeFormat,
	}, nil
}
