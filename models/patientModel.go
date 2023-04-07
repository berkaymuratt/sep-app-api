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
	Age       int     `bson:"age" json:"age"`
	Height    int     `bson:"height" json:"height"`
	Weight    float64 `bson:"weight" json:"weight"`
	Address   string  `bson:"address" json:"address"`
	Telephone string  `bson:"telephone" json:"telephone"`
}

type PatientDTO struct {
	ID          primitive.ObjectID `json:"id"`
	DoctorId    primitive.ObjectID `json:"doctor_id"`
	UserId      string             `json:"user_id"`
	PatientInfo PatientInfo        `json:"patient_info"`
	Doctor      Doctor             `json:"doctor"`
}

type GetPatientDbResponse struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	DoctorId     primitive.ObjectID `bson:"_doctor_id"`
	UserId       string             `bson:"user_id"`
	UserPassword string             `bson:"user_password"`
	PatientInfo  PatientInfo        `bson:"patient_info"`
	Doctors      []Doctor           `bson:"doctors"`
}
