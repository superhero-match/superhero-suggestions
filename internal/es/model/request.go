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

// Request holds the parameters that are received with request at suggestions endpoint.
type Request struct {
	ID                    string   `json:"id"`
	LookingForGender      int      `json:"looking_for_gender"`
	Gender                int      `json:"gender"`
	LookingForAgeMin      int      `json:"looking_for_age_min"`
	LookingForAgeMax      int      `json:"looking_for_age_max"`
	MaxDistance           int      `json:"max_distance"`
	DistanceUnit          string   `json:"distance_unit"`
	Lat                   float64  `json:"lat"`
	Lon                   float64  `json:"lon"`
	RetrievedSuperheroIDs []string `json:"retrieved_superhero_ids"`
}
