package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Doctor struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId       string             `bson:"user_id" json:"user_id"`
	UserPassword string             `bson:"user_password" json:"user_password"`
	DoctorInfo   DoctorInfo         `bson:"doctor_info" json:"doctor_info"`
}

type DoctorInfo struct {
	Name      string `bson:"name" json:"name"`
	Surname   string `bson:"surname" json:"surname"`
	Address   string `bson:"address" json:"address"`
	Telephone string `bson:"telephone" json:"telephone"`
}
