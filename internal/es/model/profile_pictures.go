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

// ProfilePicture holds data related to profile picture of a Superhero.
type ProfilePicture struct {
	ID                int64  `json:"id"`
	SuperheroID       string `json:"superhero_id"`
	ProfilePictureURL string `json:"profile_pic_url"`
	Position          int    `json:"position"`
}
