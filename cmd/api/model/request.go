package model

// Request holds the parameters that are received with request at suggestions endpoint.
type Request struct {
	ID               string  `json:"id"`
	LookingForGender int     `json:"lookingForGender"`
	Gender           int     `json:"gender"`
	LookingForAgeMin int     `json:"lookingForAgeMin"`
	LookingForAgeMax int     `json:"lookingForAgeMax"`
	MaxDistance      int     `json:"maxDistance"`
	Lat              float64 `json:"lat"`
	Lon              float64 `json:"lon"`
}
