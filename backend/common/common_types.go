package common

import "time"

type locationMeta struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

type location struct {
	Lat      float32      `json:"lat"`
	Lon      float32      `json:"lon"`
	Accuracy int          `json:"accuracy"`
	Meta     locationMeta `json:"meta"`
}

type openingHours struct {
	Open  time.Time `json:"open"`
	Close time.Time `json:"close"`
}

type RestaurantMetadata struct {
	Name         string         `json:"name"`
	Location     location       `json:"location"`
	Chain        string         `json:"chain"`
	OpeningHours []openingHours `json:"opening_hours"`
}

type mealOption struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Diets       string  `json:"diets"`
	Price       float32 `json:"price"`
}

type dayMenu struct {
	Date    time.Time    `json:"date"`
	Options []mealOption `json:"options"`
}

type Restaurant struct {
	Meta RestaurantMetadata `json:"meta"`
	Menu []dayMenu
}
