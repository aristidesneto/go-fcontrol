package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID       bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
	// CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	// UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"`
	// DeletedAt primitive.DateTime `json:"deleted_at" bson:"deleted_at"`
}

type UserResponse struct {
	ID    bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string        `json:"name,omitempty"`
	Email string        `json:"email,omitempty"`
}
