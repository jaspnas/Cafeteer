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
	Opens  time.Time `json:"opens"`
	Closes time.Time `json:"closes"`
}

type RestaurantMetadata struct {
	Name         string         `json:"name"`
	Location     location       `json:"location"`
	Chain        string         `json:"chain"`
	OpeningHours []openingHours `json:"opening_hours"`
}

type Restaurant struct {
	Meta RestaurantMetadata `json:"meta"`
}
