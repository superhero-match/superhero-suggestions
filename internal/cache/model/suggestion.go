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
