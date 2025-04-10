package repository

import "github.com/Blesssssd/microservicio-golang/models"

type VehicleRepositoryInterface interface {
	Create(vehicle *models.Vehicle) (*models.Vehicle, error)
	GetAll() ([]models.Vehicle, error)
	GetByID(id string) (*models.Vehicle, error)
	Update(id string, vehicle *models.Vehicle) error
	Delete(id string) error
}