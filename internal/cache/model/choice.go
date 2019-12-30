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
