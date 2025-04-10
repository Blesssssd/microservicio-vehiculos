package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Blesssssd/microservicio-golang/config"
	"github.com/Blesssssd/microservicio-golang/controllers"
	"github.com/Blesssssd/microservicio-golang/models"
	"github.com/Blesssssd/microservicio-golang/repository"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// setupIntegration conecta a MongoDB y retorna un controlador real
func setupIntegration(t *testing.T) *controllers.VehicleController {
	// Ya cargamos el .env desde el script de PowerShell
	client, collection := config.ConnectMongo()

	// Limpieza al finalizar test
	t.Cleanup(func() {
		client.Disconnect(context.TODO())
	})

	// Borramos todos los documentos antes de testear
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	collection.Drop(ctx)

	repo := repository.NewVehicleRepository(collection)
	return controllers.NewVehicleController(repo)
}

func TestCRUDIntegration(t *testing.T) {
	ctrl := setupIntegration(t)

	// ------------------ CREATE ------------------
	vehicle := &models.Vehicle{
		Brand:        "Ford",
		Model:        "Ranger",
		Year:         2023,
		LicensePlate: "TEST123",
	}
	body, _ := json.Marshal(vehicle)

	req := httptest.NewRequest("POST", "/vehicles", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	http.HandlerFunc(ctrl.CreateVehicle).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var createdResponse map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &createdResponse)
	data := createdResponse["data"].(map[string]interface{})
	insertedID := data["id"].(string)

	// ------------------ GET ALL ------------------
	reqGetAll := httptest.NewRequest("GET", "/vehicles", nil)
	rrGetAll := httptest.NewRecorder()
	http.HandlerFunc(ctrl.GetVehicles).ServeHTTP(rrGetAll, reqGetAll)

	assert.Equal(t, http.StatusOK, rrGetAll.Code)
	assert.Contains(t, rrGetAll.Body.String(), "Ranger")

	// ------------------ GET BY ID ------------------
	reqGet := httptest.NewRequest("GET", "/vehicles/"+insertedID, nil)
	reqGet = mux.SetURLVars(reqGet, map[string]string{"id": insertedID})
	rrGet := httptest.NewRecorder()
	http.HandlerFunc(ctrl.GetVehicleByID).ServeHTTP(rrGet, reqGet)

	assert.Equal(t, http.StatusOK, rrGet.Code)
	assert.Contains(t, rrGet.Body.String(), "Ranger")

	// ------------------ UPDATE ------------------
	updated := &models.Vehicle{
		Brand:        "Ford",
		Model:        "Raptor",
		Year:         2024,
		LicensePlate: "RPT456",
	}
	updateBody, _ := json.Marshal(updated)

	reqUpdate := httptest.NewRequest("PUT", "/vehicles/"+insertedID, bytes.NewBuffer(updateBody))
	reqUpdate.Header.Set("Content-Type", "application/json")
	reqUpdate = mux.SetURLVars(reqUpdate, map[string]string{"id": insertedID})
	rrUpdate := httptest.NewRecorder()
	http.HandlerFunc(ctrl.UpdateVehicle).ServeHTTP(rrUpdate, reqUpdate)

	assert.Equal(t, http.StatusOK, rrUpdate.Code)
	assert.Contains(t, rrUpdate.Body.String(), "Vehicle updated successfully")

	// ------------------ DELETE ------------------
	reqDelete := httptest.NewRequest("DELETE", "/vehicles/"+insertedID, nil)
	reqDelete = mux.SetURLVars(reqDelete, map[string]string{"id": insertedID})
	rrDelete := httptest.NewRecorder()
	http.HandlerFunc(ctrl.DeleteVehicle).ServeHTTP(rrDelete, reqDelete)

	assert.Equal(t, http.StatusOK, rrDelete.Code)
	assert.Contains(t, rrDelete.Body.String(), "Vehicle deleted successfully")
}