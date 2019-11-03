package model

type Superhero struct {
	ID                    string  `json:"id"`
	LookingForGender      int     `json:"lookingForGender"`
	Age                   int     `json:"age"`
	LookingForAgeMin      int     `json:"lookingForAgeMin"`
	LookingForAgeMax      int     `json:"lookingForAgeMax"`
	LookingForDistanceMax int     `json:"lookingForDistanceMax"`
	DistanceUnit          string  `json:"distanceUnit"`
	Lat                   float64 `json:"lat"`
	Lon                   float64 `json:"lon"`
	Country               string  `json:"country"`
	City                  string  `json:"city"`
	SuperPower            string  `json:"superPower"`
	AccountType           string  `json:"accountType"`
}
