package modals

import (
	"gopkg.in/mgo.v2/bson"
)

type user struct {
	id     bson.ObjectId
	name   string
	gender string
	age    int
}
