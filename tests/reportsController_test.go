package tests

import (
	"fmt"
	"github.com/berkaymuratt/sep-app-api/controllers"
	"github.com/berkaymuratt/sep-app-api/dtos"
	"github.com/berkaymuratt/sep-app-api/mocks/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http/httptest"
	"testing"
)

var reportsController controllers.ReportsController

func TestReportsController_GetReports(t *testing.T) {
	controller := gomock.NewController(t)

	mockReportsService := services.NewMockReportsServiceI(controller)
	reportsController = controllers.NewReportsController(mockReportsService)

	router := fiber.New()
	router.Get("/api/reports", reportsController.GetReports)

	doctorId := primitive.NewObjectID()

	reports := []*dtos.ReportDto{
		{},
	}

	mockReportsService.EXPECT().GetReports(&doctorId, nil).Return(reports, nil)

	target := fmt.Sprintf("/api/reports?doctor_id=%s", doctorId.Hex())
	req := httptest.NewRequest("GET", target, nil)
	response, _ := router.Test(req)

	assert.Equal(t, 200, response.StatusCode)
}

func TestReportsController_GetReportById(t *testing.T) {
	controller := gomock.NewController(t)

	mockReportsService := services.NewMockReportsServiceI(controller)
	reportsController = controllers.NewReportsController(mockReportsService)

	router := fiber.New()
	router.Get("/api/reports/:id", reportsController.GetReportById)

	report := dtos.ReportDto{
		ID: primitive.NewObjectID(),
	}

	mockReportsService.EXPECT().GetReportById(report.ID).Return(&report, nil)

	target := fmt.Sprintf("/api/reports/%s", report.ID.Hex())
	req := httptest.NewRequest("GET", target, nil)
	response, _ := router.Test(req)

	assert.Equal(t, 200, response.StatusCode)
}
