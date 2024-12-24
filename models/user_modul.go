package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModul struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	User_ID  primitive.ObjectID `bson:"user_id" json:"user_id"`
	Modul_ID primitive.ObjectID `bson:"modul_id" json:"modul_id"`
}

