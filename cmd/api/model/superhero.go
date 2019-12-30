package model

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
	HasLikedMe        bool             `json:"hasLikedMe"`
	CreatedAt         string           `json:"createdAt"`
}
