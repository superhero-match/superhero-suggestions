package model

import "encoding/json"

// ProfilePicture holds data related to profile picture of a Superhero.
type ProfilePicture struct {
	ID                int64  `json:"id"`
	SuperheroID       string `json:"superheroId"`
	ProfilePictureURL string `json:"profilePicUrl"`
	Position          int    `json:"position"`
}

// MarshalBinary ...
func (p *ProfilePicture) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

// UnmarshalBinary ...
func (p ProfilePicture) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &p)
}