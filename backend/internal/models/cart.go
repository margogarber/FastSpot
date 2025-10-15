package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Cart represents a shopping cart
type Cart struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    string             `bson:"userId,omitempty" json:"userId,omitempty"`
	SessionID string             `bson:"sessionId,omitempty" json:"sessionId,omitempty"`
	Items     []CartItem         `bson:"items" json:"items"`
	TotalUSD  float64            `bson:"totalUSD" json:"totalUSD"`
	Currency  string             `bson:"currency" json:"currency"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// CartItem represents an item in the cart
type CartItem struct {
	ProductID         primitive.ObjectID `bson:"productId" json:"productId"`
	Name              string             `bson:"name" json:"name"`
	Image             string             `bson:"image" json:"image"`
	Qty               int                `bson:"qty" json:"qty"`
	UnitPriceUSD      float64            `bson:"unitPriceUSD" json:"unitPriceUSD"`
	TotalUSD          float64            `bson:"totalUSD" json:"totalUSD"`
	ChosenIngredients []string           `bson:"chosenIngredients" json:"chosenIngredients"`
	ChosenOptions     map[string]string  `bson:"chosenOptions" json:"chosenOptions"`
}
