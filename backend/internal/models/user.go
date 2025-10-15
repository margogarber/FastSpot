package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Role         string             `bson:"role" json:"role"` // "admin" | "guest"
	Name         string             `bson:"name" json:"name"`
	Email        string             `bson:"email" json:"email"`
	Phone        string             `bson:"phone" json:"phone"`
	PasswordHash string             `bson:"passwordHash" json:"-"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone"`
	Password string `json:"password" binding:"required,min=6"`
}
