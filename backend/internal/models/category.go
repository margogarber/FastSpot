package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Category represents a product category
type Category struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Slug      string             `bson:"slug" json:"slug"`
	Image     string             `bson:"image" json:"image"`
	IsActive  bool               `bson:"isActive" json:"isActive"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
}
