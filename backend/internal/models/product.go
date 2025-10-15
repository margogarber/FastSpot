package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CategoryID  primitive.ObjectID `bson:"categoryId" json:"categoryId"`
	Name        string             `bson:"name" json:"name"`
	Slug        string             `bson:"slug" json:"slug"`
	Description string             `bson:"description" json:"description"`
	PriceUSD    float64            `bson:"priceUSD" json:"priceUSD"`
	Image       string             `bson:"image" json:"image"`
	IsActive    bool               `bson:"isActive" json:"isActive"`
	Ingredients []Ingredient       `bson:"ingredients" json:"ingredients"`
	Options     []ProductOption    `bson:"options" json:"options"`
	Tags        []string           `bson:"tags" json:"tags"`
	CreatedAt   time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
}

type Ingredient struct {
	Key             string `bson:"key" json:"key"`
	Label           string `bson:"label" json:"label"`
	DefaultIncluded bool   `bson:"defaultIncluded" json:"defaultIncluded"`
}

type ProductOption struct {
	Key     string         `bson:"key" json:"key"`
	Label   string         `bson:"label" json:"label"`
	Type    string         `bson:"type" json:"type"` // "single" | "multiple"
	Choices []OptionChoice `bson:"choices" json:"choices"`
}

type OptionChoice struct {
	Value         string  `bson:"value" json:"value"`
	Label         string  `bson:"label" json:"label"`
	ExtraPriceUSD float64 `bson:"extraPriceUSD" json:"extraPriceUSD"`
}

type CreateProductRequest struct {
	CategoryID  string          `json:"categoryId" binding:"required"`
	Name        string          `json:"name" binding:"required"`
	Slug        string          `json:"slug" binding:"required"`
	Description string          `json:"description"`
	PriceUSD    float64         `json:"priceUSD" binding:"required,gt=0"`
	Image       string          `json:"image"`
	IsActive    bool            `json:"isActive"`
	Ingredients []Ingredient    `json:"ingredients"`
	Options     []ProductOption `json:"options"`
	Tags        []string        `json:"tags"`
}
