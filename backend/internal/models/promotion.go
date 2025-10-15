package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Promotion represents a promotional offer
type Promotion struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	StartsAt    time.Time          `bson:"startsAt" json:"startsAt"`
	EndsAt      time.Time          `bson:"endsAt" json:"endsAt"`
	BannerImage string             `bson:"bannerImage" json:"bannerImage"`
	IsActive    bool               `bson:"isActive" json:"isActive"`
	AppliesTo   []string           `bson:"appliesTo" json:"appliesTo"` // Array of product slugs
	CreatedAt   time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
}
