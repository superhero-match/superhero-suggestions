package model

// Request holds the parameters that are received with request at suggestions endpoint.
type Request struct {
	ID               string   `json:"id"`
	LookingForGender int      `json:"looking_for_gender"`
	Gender           int      `json:"gender"`
	LookingForAgeMin int      `json:"looking_for_age_min"`
	LookingForAgeMax int      `json:"looking_for_age_max"`
	MaxDistance      int      `json:"max_distance"`
	DistanceUnit     string   `json:"distance_unit"`
	Lat              float64  `json:"lat"`
	Lon              float64  `json:"lon"`
	SuperheroIDs     []string `json:"superhero_ids"`
}
