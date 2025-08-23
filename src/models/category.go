package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Category struct {
	ID     bson.ObjectID `form:"id"   bson:"_id,omitempty"  json:"id,omitempty"`
	UserId bson.ObjectID `form:"user_id" bson:"user_id,omitempty" json:"user_id,omitempty"`
	Name   string        `form:"name" bson:"name,omitempty" json:"name,omitempty"`
	Color  string        `form:"color" bson:"color,omitempty" json:"color,omitempty"`
	Type   string        `form:"type" bson:"type,omitempty" json:"type,omitempty"`
}
