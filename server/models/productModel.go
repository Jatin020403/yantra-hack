package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID            primitive.ObjectID `bson:"_id"`
	ProductName   *string            `json:"product_name" validate:"required,min=2,max=100"`
	ProductOrigin *string            `json:"product_origin" validate:"required,min=6"`
	Phone         *string            `json:"phone" validate:"required"`
	Token         *string            `json:"token"`
	RefreshToken  *string            `json:"refresh_token"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
	UserID        string             `json:"user_id"`
	Locations     []Location         `json:"locations"`
	Damaged       bool               `json:"damaged"`
	DateOfLeaving time.Time          `json:"date_of_leaving"`
	Ratings       []Rating           `json:"ratings"`
}

type Location struct {
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timestamp time.Time `json:"timestamp"`
}

type Rating struct {
	UserID      string    `json:"user_id"`
	Rating      float64   `json:"rating"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
