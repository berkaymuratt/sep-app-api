package services

import "github.com/berkaymuratt/sep-app-api/models"

type BodyPartsService struct{}

func NewBodyPartsService() BodyPartsService {
	return BodyPartsService{}
}

func (service BodyPartsService) GetBodyParts() ([]*models.BodyPart, error) {
	// TODO: implement function
	return nil, nil
}
