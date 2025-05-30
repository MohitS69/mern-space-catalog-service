package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// PriceConfiguration represents the price configuration for a category
type PriceConfiguration struct {
	PriceType        string   `bson:"priceType" json:"priceType" binding:"required,oneof=base aditional"`
	AvailableOptions []string `bson:"availableOptions" json:"availableOptions" binding:"required,min=1"`
}

// Attribute represents an attribute of a category
type Attribute struct {
	Name             string      `bson:"name" json:"name" binding:"required"`
	WidgetType       string      `bson:"widgetType" json:"widgetType" binding:"required,oneof=switch radio"`
	DefaultValue     interface{} `bson:"defaultValue" json:"defaultValue" binding:"required"`
	AvailableOptions []string    `bson:"availableOptions" json:"availableOptions" binding:"required,min=1"`
}

// Category represents the main category document
type Category struct {
	ID                 primitive.ObjectID            `bson:"_id,omitempty" json:"id,omitempty"`
	Name               string                        `bson:"name" json:"name" binding:"required"`
	PriceConfiguration map[string]PriceConfiguration `bson:"priceConfiguration" json:"priceConfiguration" binding:"required"`
	Attributes         []Attribute                   `bson:"attributes" json:"attributes" binding:"required,min=1"`
	CreatedAt          time.Time                     `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt          time.Time                     `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

func GetCategoryCollection(db *mongo.Database) *mongo.Collection {
	return db.Collection("category")
}
