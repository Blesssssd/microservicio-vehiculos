package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/Blesssssd/microservicio-golang/controllers"
	 mockrepo "github.com/Blesssssd/microservicio-golang/controllers_test"
	"github.com/Blesssssd/microservicio-golang/models"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupController(t *testing.T) (*controllers.VehicleController, *mockrepo.MockVehicleRepositoryInterface) {
	ctrl := gomock.NewController(t)
	mockRepo := mockrepo.NewMockVehicleRepositoryInterface(ctrl)
	controller := &controllers.VehicleController{Repo: mockRepo}
	return controller, mockRepo
}


func TestCreateVehicle(t *testing.T) {
    controller, mockRepo := setupController(t)

    vehicle := &models.Vehicle{
        Brand:        "Toyota",
        Model:        "Corolla",
        Year:         2020,
        LicensePlate: "ABC123",
    }

    mockRepo.EXPECT().Create(gomock.Any()).Return(vehicle, nil)

    body, _ := json.Marshal(vehicle)
    req := httptest.NewRequest("POST", "/vehicles", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(controller.CreateVehicle)
    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusCreated, rr.Code)
    assert.Contains(t, rr.Body.String(), "Vehicle created successfully")
}

func TestGetVehicles(t *testing.T) {
	controller, mockRepo := setupController(t)

	mockVehicles := []models.Vehicle{
		{Brand: "Toyota", Model: "Yaris", Year: 2020, LicensePlate: "ABC123"},
		{Brand: "Mazda", Model: "3", Year: 2019, LicensePlate: "XYZ789"},
	}

	mockRepo.EXPECT().GetAll().Return(mockVehicles, nil)

	req := httptest.NewRequest("GET", "/vehicles", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controller.GetVehicles)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "List of vehicles")
}


func TestGetVehicleByID(t *testing.T) {
	controller, mockRepo := setupController(t)

	mockVehicle := &models.Vehicle{
		Brand: "Chevrolet", Model: "Onix", Year: 2021, LicensePlate: "ONX123",
	}

	mockRepo.EXPECT().GetByID("123").Return(mockVehicle, nil)

	req := httptest.NewRequest("GET", "/vehicles/123", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "123"})
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controller.GetVehicleByID)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Vehicle found")
}


func TestUpdateVehicle(t *testing.T) {
	controller, mockRepo := setupController(t)

	updatedVehicle := &models.Vehicle{
		Brand: "Hyundai", Model: "Tucson", Year: 2022, LicensePlate: "HYD456",
	}

	mockRepo.EXPECT().Update("123", gomock.Any()).Return(nil)

	body, _ := json.Marshal(updatedVehicle)
	req := httptest.NewRequest("PUT", "/vehicles/123", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req = mux.SetURLVars(req, map[string]string{"id": "123"})
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controller.UpdateVehicle)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Vehicle updated successfully")
}

func TestDeleteVehicle(t *testing.T) {
	controller, mockRepo := setupController(t)

	mockRepo.EXPECT().Delete("123").Return(nil)

	req := httptest.NewRequest("DELETE", "/vehicles/123", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "123"})
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controller.DeleteVehicle)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Vehicle deleted successfully")
}
