package structType

import (
	"time"
)

/* type Item struct {
	Id int `json:"id" bson:"id"`
	Title string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Price float64 `json:"price" bson:"price"`
	Stock int	`json:"stock" bson:"stock"`
	Category string	`json:"category" bson:"category"`
	Thumbnail string `json:"thumbnail" bson:"thumbnail"`
	Images []string `json:"images" bson:"images"`
} */

type Item struct{
	Id int `bson:"id"`
	Type string `json:"type" bson:"type"`
	LastAccess time.Time
	Status string
	Geometry GeoData 
}

type GeoData struct{
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
	Type string `json:"type" bson:"type"`
}