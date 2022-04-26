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

package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	err := os.Setenv("TEST_CONFIG", "config.test.yml")
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	// App
	assert.Equal(t, ":4000", cfg.App.Port, "The port should be :4000.")
	assert.Equal(t, "2006-01-02T15:04:05", cfg.App.TimeFormat, "The time format should be 2006-01-02T15:04:05.")
	assert.Equal(t, 10, cfg.App.PageSize, "The page size should be 10.")

	// Elasticsearch
	assert.Equal(t, "localhost", cfg.ES.Host, "The host should be localhost.")
	assert.Equal(t, "9200", cfg.ES.Port, "The port should be 9200.")
	assert.Equal(t, "superheromatch", cfg.ES.Cluster, "The cluster should be superheromatch.")
	assert.Equal(t, "superhero", cfg.ES.Index, "The index should be superhero.")
	assert.Equal(t, 50, cfg.ES.BatchSize, "The batch size should be 50.")

	// Cache
	assert.Equal(t, "localhost", cfg.Cache.Address, "The address should be localhost.")
	assert.Equal(t, ":6379", cfg.Cache.Port, "The port should be :6379.")
	assert.Equal(t, "Awesome85**", cfg.Cache.Password, "The password should be Awesome85**.")
	assert.Equal(t, 0, cfg.Cache.DB, "The db should be 0.")
	assert.Equal(t, 25, cfg.Cache.PoolSize, "The pool size should be 25.")
	assert.Equal(t, 10, cfg.Cache.MinimumIdleConnections, "The minimum idle connections should be 10.")
	assert.Equal(t, 1, cfg.Cache.MaximumRetries, "The maximum retries should be 1.")
	assert.Equal(t, "suggestion.%s", cfg.Cache.SuggestionKeyFormat, "The suggestion key format should be suggestion.%s.")
	assert.Equal(t, "choice.%s.%s", cfg.Cache.ChoiceKeyFormat, "The choice key format should be choice.%s.%s.")
	assert.Equal(t, "likes.for.%s", cfg.Cache.LikesKeyFormat, "The likes key format should be likes.for.%s.")
}
