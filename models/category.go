package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PriceConfiguration represents the price configuration for a category
type PriceConfiguration struct {
	PriceType        string   `bson:"priceType" json:"priceType" validate:"required,oneof=base aditional"`
	AvailableOptions []string `bson:"availableOptions" json:"availableOptions" validate:"required,min=1"`
}

// Attribute represents an attribute of a category
type Attribute struct {
	Name             string      `bson:"name" json:"name" validate:"required"`
	WidgetType       string      `bson:"widgetType" json:"widgetType" validate:"required,oneof=switch radio"`
	DefaultValue     interface{} `bson:"defaultValue" json:"defaultValue" validate:"required"`
	AvailableOptions []string    `bson:"availableOptions" json:"availableOptions" validate:"required,min=1"`
}

// Category represents the main category document
type Category struct {
	ID                 primitive.ObjectID            `bson:"_id,omitempty" json:"id,omitempty"`
	Name               string                        `bson:"name" json:"name" validate:"required"`
	PriceConfiguration map[string]PriceConfiguration `bson:"priceConfiguration" json:"priceConfiguration" validate:"required"`
	Attributes         []Attribute                   `bson:"attributes" json:"attributes" validate:"required,min=1"`
	CreatedAt          time.Time                     `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt          time.Time                     `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

// CategoryRepository handles database operations for categories
type CategoryRepository struct {
	collection *mongo.Collection
}

// NewCategoryRepository creates a new category repository
func NewCategoryRepository(db *mongo.Database) *CategoryRepository {
	return &CategoryRepository{
		collection: db.Collection("categories"),
	}
}

