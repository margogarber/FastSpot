package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Order represents a customer order
type Order struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID       string             `bson:"userId,omitempty" json:"userId,omitempty"`
	SessionID    string             `bson:"sessionId,omitempty" json:"sessionId,omitempty"`
	OrderNumber  string             `bson:"orderNumber" json:"orderNumber"`
	Items        []OrderItem        `bson:"items" json:"items"`
	TotalUSD     float64            `bson:"totalUSD" json:"totalUSD"`
	Currency     string             `bson:"currency" json:"currency"`
	Status       string             `bson:"status" json:"status"` // new, confirmed, preparing, ready, delivering, completed, cancelled
	Payment      Payment            `bson:"payment" json:"payment"`
	Delivery     Delivery           `bson:"delivery" json:"delivery"`
	CustomerInfo CustomerInfo       `bson:"customerInfo" json:"customerInfo"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	ProductID         primitive.ObjectID `bson:"productId" json:"productId"`
	Name              string             `bson:"name" json:"name"`
	Image             string             `bson:"image" json:"image"`
	Qty               int                `bson:"qty" json:"qty"`
	UnitPriceUSD      float64            `bson:"unitPriceUSD" json:"unitPriceUSD"`
	TotalUSD          float64            `bson:"totalUSD" json:"totalUSD"`
	ChosenIngredients []string           `bson:"chosenIngredients" json:"chosenIngredients"`
	ChosenOptions     map[string]string  `bson:"chosenOptions" json:"chosenOptions"`
}

// Payment represents payment information
type Payment struct {
	Method string `bson:"method" json:"method"` // card, applepay, googlepay, cash
	Status string `bson:"status" json:"status"` // pending, completed, failed
	TxnID  string `bson:"txnId,omitempty" json:"txnId,omitempty"`
}

// Delivery represents delivery information
type Delivery struct {
	Type     string          `bson:"type" json:"type"` // pickup, delivery
	Address  DeliveryAddress `bson:"address,omitempty" json:"address,omitempty"`
	ETA      time.Time       `bson:"eta,omitempty" json:"eta,omitempty"`
	Tracking []TrackingEvent `bson:"tracking" json:"tracking"`
}

// DeliveryAddress represents delivery address
type DeliveryAddress struct {
	Street  string `bson:"street" json:"street"`
	City    string `bson:"city" json:"city"`
	ZipCode string `bson:"zipCode" json:"zipCode"`
	Notes   string `bson:"notes,omitempty" json:"notes,omitempty"`
}

// TrackingEvent represents an order tracking event
type TrackingEvent struct {
	Timestamp time.Time `bson:"ts" json:"ts"`
	Status    string    `bson:"status" json:"status"`
	Note      string    `bson:"note,omitempty" json:"note,omitempty"`
}

// CustomerInfo represents customer contact information
type CustomerInfo struct {
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email,omitempty" json:"email,omitempty"`
	Phone string `bson:"phone" json:"phone"`
}
