/*
  Copyright (C) 2019 - 2021 MWSOFT
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

// Cache holds all the configuration settings for the Redis client.
type Cache struct {
	Address                string `env:"REDIS_ADDRESS" yaml:"address" default:"192.168.0.64"`
	Port                   string `env:"REDIS_PORT" yaml:"port" default:":6379"`
	Password               string `env:"REDIS_PASSWORD" yaml:"password" default:"Awesome85**"`
	DB                     int    `env:"REDIS_DB" yaml:"db" default:"0"`
	PoolSize               int    `env:"REDIS_POOL_SIZE" yaml:"pool_size" default:"25"`
	MinimumIdleConnections int    `env:"REDIS_MINIMUM_IDLE_CONNECTIONS" yaml:"minimum_idle_connections" default:"10"`
	MaximumRetries         int    `env:"REDIS_MAXIMUM_RETRIES" yaml:"maximum_retries" default:"1"`
	SuggestionKeyFormat    string `env:"SUPERHERO_SUGGESTIONS_REDIS_SUGGESTION_KEY_FORMAT" yaml:"suggestion_key_format" default:"suggestion.%s"`
	ChoiceKeyFormat        string `env:"SUPERHERO_SUGGESTIONS_REDIS_CHOICE_KEY_FORMAT" yaml:"choice_key_format" default:"choice.%s.%s"`
	LikesKeyFormat         string `env:"SUPERHERO_SUGGESTIONS_REDIS_LIKES_KEY_FORMAT" yaml:"likes_key_format" default:"likes.for.%s"`
}
