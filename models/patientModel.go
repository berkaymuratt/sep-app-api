package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Patient struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	DoctorId     primitive.ObjectID `bson:"_doctor_id" json:"doctor_id"`
	UserId       string             `bson:"user_id" json:"user_id"`
	UserPassword string             `bson:"user_password" json:"user_password"`
	PatientInfo  PatientInfo        `bson:"patient_info" json:"patient_info"`
}

type PatientInfo struct {
	Name      string  `bson:"name" json:"name"`
	Surname   string  `bson:"surname" json:"surname"`
	Gender    string  `bson:"gender" json:"gender"`
	Age       int     `bson:"age" json:"age"`
	Height    int     `bson:"height" json:"height"`
	Weight    float64 `bson:"weight" json:"weight"`
	Address   string  `bson:"address" json:"address"`
	Telephone string  `bson:"telephone" json:"telephone"`
}
