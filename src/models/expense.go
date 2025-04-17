package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Expense struct {
	ID         bson.ObjectID `form:"id" bson:"_id,omitempty" json:"id"`
	Name       string        `form:"name" bson:"name" json:"name"`
	CategoryID bson.ObjectID `form:"category_id" bson:"category_id,omitempty" json:"category_id"`
}
