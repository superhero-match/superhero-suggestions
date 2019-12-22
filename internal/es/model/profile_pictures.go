package model

// ProfilePicture holds data related to profile picture of a Superhero.
type ProfilePicture struct {
	ID                int64  `json:"id"`
	SuperheroID       string `json:"superhero_id"`
	ProfilePictureURL string `json:"profile_pic_url"`
	Position          int    `json:"position"`
}
