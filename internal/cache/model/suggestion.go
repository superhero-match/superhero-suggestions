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
package model

import "encoding/json"

// Superhero holds the user profile data returned as suggestion.
type Superhero struct {
	ID                string           `json:"id"`
	SuperheroName     string           `json:"superheroName"`
	MainProfilePicURL string           `json:"mainProfilePicUrl"`
	ProfilePictures   []ProfilePicture `json:"profilePictures"`
	Gender            int              `json:"gender"`
	Age               int              `json:"age"`
	Lat               float64          `json:"lat"`
	Lon               float64          `json:"lon"`
	Birthday          string           `json:"birthday"`
	Country           string           `json:"country"`
	City              string           `json:"city"`
	SuperPower        string           `json:"superpower"`
	AccountType       string           `json:"accountType"`
	CreatedAt         string           `json:"createdAt"`
}

// MarshalBinary ...
func (s Superhero) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

// UnmarshalBinary ...
func (s *Superhero) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &s)
}
