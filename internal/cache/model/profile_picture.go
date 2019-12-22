package model

// ProfilePicture holds data related to profile picture of a Superhero.
type ProfilePicture struct {
	ID                int64  `json:"id"`
	SuperheroID       string `json:"superheroId"`
	ProfilePictureURL string `json:"profilePicUrl"`
	Position          int    `json:"position"`
}
