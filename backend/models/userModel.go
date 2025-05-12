package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	First_name    *string            `bson:"first_name" json:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string            `bson:"last_name" json:"last_name" validate:"min=2,max=100"`
	Password      *string            `bson:"password" json:"password" validate:"required,min=6"`
	Email         *string            `bson:"email" json:"email" validate:"required,email"`
	Phone         *string            `bson:"phone" json:"phone"`
	Token         *string            `bson:"token" json:"token"`
	User_type     *string            `bson:"user_type" json:"user_type" validate:"required,oneof=ADMIN USER"`
	Refresh_token *string            `bson:"refresh_token" json:"refresh_token"`
	Created_at    time.Time          `bson:"created_at" json:"created_at"`
	Updated_at    time.Time          `bson:"updated_at" json:"updated_at"`
	User_id       string             `bson:"user_id" json:"user_id"`
	Provider      string             `bson:"provider" json:"provider" validate:"required,oneof=Google Local"`
	ImageURL      string             `bson:"image_url" json:"image_url"`
}
