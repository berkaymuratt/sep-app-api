package tests

import (
	"fmt"
	"github.com/berkaymuratt/sep-app-api/controllers"
	"github.com/berkaymuratt/sep-app-api/mocks/services"
	"github.com/berkaymuratt/sep-app-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http/httptest"
	"testing"
)

var bodyPartsController controllers.BodyPartsController

func TestBodyPartsController_GetBodyParts(t *testing.T) {
	controller := gomock.NewController(t)

	mockBodyPartsService := services.NewMockBodyPartsServiceI(controller)
	bodyPartsController = controllers.NewBodyPartsController(mockBodyPartsService)

	router := fiber.New()
	router.Get("/api/body-parts", bodyPartsController.GetBodyParts)

	bodyParts := []*models.BodyPart{
		{
			ID:   primitive.NewObjectID(),
			Name: "Body Part",
		},
	}

	mockBodyPartsService.EXPECT().GetBodyParts().Return(bodyParts, nil)

	target := fmt.Sprintf("/api/body-parts")
	req := httptest.NewRequest("GET", target, nil)
	response, _ := router.Test(req)

	assert.Equal(t, 200, response.StatusCode)
}
