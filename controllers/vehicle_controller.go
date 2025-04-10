package controllers

import (
    "encoding/json"
    "github.com/Blesssssd/microservicio-golang/models"
    "github.com/Blesssssd/microservicio-golang/repository"
    "net/http"

    "github.com/gorilla/mux"
)

type VehicleController struct {
    Repo repository.VehicleRepositoryInterface
}

func NewVehicleController(repo repository.VehicleRepositoryInterface) *VehicleController {
    return &VehicleController{Repo: repo}
}

func (vc *VehicleController) CreateVehicle(w http.ResponseWriter, r *http.Request) {
    var vehicle models.Vehicle
    json.NewDecoder(r.Body).Decode(&vehicle)

    result, err := vc.Repo.Create(&vehicle)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Vehicle created successfully",
        "data":    result,
    })
}

func (vc *VehicleController) GetVehicles(w http.ResponseWriter, r *http.Request) {
    vehicles, err := vc.Repo.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message":  "List of vehicles",
        "vehicles": vehicles,
    })
}

func (vc *VehicleController) GetVehicleByID(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    vehicle, err := vc.Repo.GetByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Vehicle found",
        "vehicle": vehicle,
    })
}

func (vc *VehicleController) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    var vehicle models.Vehicle
    json.NewDecoder(r.Body).Decode(&vehicle)

    err := vc.Repo.Update(id, &vehicle)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Vehicle updated successfully",
    })
}

func (vc *VehicleController) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]

    err := vc.Repo.Delete(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Vehicle deleted successfully",
    })
}
