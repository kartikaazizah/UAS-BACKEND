package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username   string             `bson:"username" json:"username"`
	Email      string             `bson:"email" json:"email"`
	Role       primitive.ObjectID `bson:"role" json:"role"`
	Jenis_User primitive.ObjectID `bson:"jenis_user" json:"jenis_user"`
	Auth_Key   string             `bson:"auth_key" json:"auth_key"`
	Token      string             `bson:"token" json:"token"`
}
