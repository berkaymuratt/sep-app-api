package dtos

import (
	"github.com/berkaymuratt/sep-app-api/dbdtos"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DiseaseDto struct {
	ID       primitive.ObjectID `json:"id"`
	Symptoms []*SymptomDto      `json:"symptoms,omitempty"`
	Name     string             `json:"name"`
	Details  string             `json:"details"`
}

func DiseaseDtoFromDiseaseDbResponse(diseaseData *dbdtos.GetDiseaseDbResponse, symptomsDtos []*SymptomDto) (*DiseaseDto, error) {
	diseaseDto := DiseaseDto{
		ID:       diseaseData.ID,
		Symptoms: symptomsDtos,
		Name:     diseaseData.Name,
		Details:  diseaseData.Details,
	}

	return &diseaseDto, nil
}
