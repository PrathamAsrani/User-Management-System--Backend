package modals

import "gopkg.in/mgo.v2/bson"

type UserModal struct {
	Id    bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name  string        `json:"name"`
	Email string        `json:"email"`
}
