package model

import "gopkg.in/olivere/elastic.v7"

// Superhero holds the user profile data returned as suggestion.
type Superhero struct {
	ID                    string           `json:"superhero_id"`
	Email                 string           `json:"email"`
	Name                  string           `json:"name"`
	SuperheroName         string           `json:"superhero_name"`
	MainProfilePicURL     string           `json:"main_profile_pic_url"`
	ProfilePictures       []ProfilePicture `json:"profile_pictures"`
	Gender                int              `json:"gender"`
	LookingForGender      int              `json:"looking_for_gender"`
	Age                   int              `json:"age"`
	LookingForAgeMin      int              `json:"looking_for_age_min"`
	LookingForAgeMax      int              `json:"looking_for_age_max"`
	LookingForDistanceMax int              `json:"looking_for_distance_max"`
	DistanceUnit          string           `json:"distance_unit"`
	Location              elastic.GeoPoint `json:"location"`
	Birthday              string           `json:"birthday"`
	Country               string           `json:"country"`
	City                  string           `json:"city"`
	SuperPower            string           `json:"superpower"`
	AccountType           string           `json:"account_type"`
}
