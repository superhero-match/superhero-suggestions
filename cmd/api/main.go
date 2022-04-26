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
package main

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"

	"github.com/superhero-match/superhero-suggestions/cmd/api/controller"
	"github.com/superhero-match/superhero-suggestions/cmd/api/service"
	"github.com/superhero-match/superhero-suggestions/internal/cache"
	"github.com/superhero-match/superhero-suggestions/internal/config"
	"github.com/superhero-match/superhero-suggestions/internal/es"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	esClient, err := elastic.NewClient(
		elastic.SetURL(
			fmt.Sprintf(
				"http://%s:%s",
				cfg.ES.Host,
				cfg.ES.Port,
			),
		),
	)
	if err != nil {
		panic(err)
	}

	e := es.New(esClient, cfg.ES.Index, cfg.ES.BatchSize)

	redisClient := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s%s", cfg.Cache.Address, cfg.Cache.Port),
		Password:     cfg.Cache.Password,
		DB:           cfg.Cache.DB,
		PoolSize:     cfg.Cache.PoolSize,
		MinIdleConns: cfg.Cache.MinimumIdleConnections,
		MaxRetries:   cfg.Cache.MaximumRetries,
	})

	_, err = redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}

	c := cache.New(redisClient, cfg.Cache.LikesKeyFormat, cfg.Cache.SuggestionKeyFormat)

	srv := service.New(e, c, cfg.App.PageSize, cfg.Cache.ChoiceKeyFormat)

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	ctrl := controller.New(srv, logger, cfg.App.TimeFormat)
	r := ctrl.RegisterRoutes()
	err = r.Run(cfg.App.Port)
	if err != nil {
		panic(err)
	}
}
