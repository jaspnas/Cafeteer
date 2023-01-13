package common

import "time"

type locationMeta struct {
	Country string `json:"country" bson:"country"`
	City    string `json:"city" bson:"city"`
}

type location struct {
	Lat      float32      `json:"lat" bson:"lat"`
	Lon      float32      `json:"lon" bson:"lon"`
	Accuracy int          `json:"accuracy" bson:"accuracy"`
	Meta     locationMeta `json:"meta" bson:"meta"`
}

type openingHours struct {
	Open  time.Time `json:"open" bson:"open"`
	Close time.Time `json:"close" bson:"close"`
}

type RestaurantMetadata struct {
	Name         string         `json:"name" bson:"name"`
	Location     location       `json:"location" bson:"location"`
	Chain        string         `json:"chain" bson:"chain"`
	OpeningHours []openingHours `json:"opening_hours" bson:"opening_hours"`
}

type mealOption struct {
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Diets       string  `json:"diets" bson:"diets"`
	Price       float32 `json:"price" bson:"price"`
}

type dayMenu struct {
	Date    time.Time    `json:"date" bson:"date"`
	Options []mealOption `json:"options" bson:"options"`
}

type Restaurant struct {
	Meta RestaurantMetadata `json:"meta" bson:"meta"`
	Menu []dayMenu          `json:"menu" bson:"menu"`
}
