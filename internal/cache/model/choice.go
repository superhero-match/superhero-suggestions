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
package model

import (
	"encoding/json"
)

// Choice holds the information about the choice that a user made, e.g. like or dislike.
type Choice struct {
	ID                string `json:"id"`
	Choice            int64  `json:"choice"`
	SuperheroID       string `json:"superheroID"`
	ChosenSuperheroID string `json:"chosenSuperheroID"`
	CreatedAt         string `json:"createdAt"`
}

// MarshalBinary ...
func (c Choice) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

// UnmarshalBinary ...
func (c *Choice) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &c)
}
