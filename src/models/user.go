package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID       bson.ObjectID `form:"id" bson:"_id,omitempty" json:"id"`
	Name     string        `form:"name" bson:"name" json:"name"`
	Email    string        `form:"email" bson:"email" json:"email"`
	Password string        `form:"password" bson:"password" json:"password"`
}

type UserResponse struct {
	ID    bson.ObjectID `form:"id" bson:"_id,omitempty" json:"id,omitempty"`
	Name  string        `form:"name" bson:"name" json:"name,omitempty"`
	Email string        `form:"email" bson:"email" json:"email,omitempty"`
}
