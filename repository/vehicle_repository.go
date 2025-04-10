package repository

import (
    "context"
    "github.com/Blesssssd/microservicio-golang/models"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type VehicleRepository struct {
    Collection *mongo.Collection
}

func NewVehicleRepository(col *mongo.Collection) *VehicleRepository {
    return &VehicleRepository{Collection: col}
}

func (r *VehicleRepository) Create(vehicle *models.Vehicle) (*models.Vehicle, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.Collection.InsertOne(ctx, vehicle)
	if err != nil {
		return nil, err
	}

	vehicle.ID = result.InsertedID.(primitive.ObjectID) // asegúrate de que ID esté en el modelo
	return vehicle, nil
}


func (r *VehicleRepository) GetAll() ([]models.Vehicle, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := r.Collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }

    var vehicles []models.Vehicle
    if err = cursor.All(ctx, &vehicles); err != nil {
        return nil, err
    }

    return vehicles, nil
}

func (r *VehicleRepository) GetByID(id string) (*models.Vehicle, error) {
    objID, _ := primitive.ObjectIDFromHex(id)
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var vehicle models.Vehicle
    err := r.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&vehicle)
    return &vehicle, err
}

func (r *VehicleRepository) Update(id string, vehicle *models.Vehicle) error {
    objID, _ := primitive.ObjectIDFromHex(id)
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := r.Collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": vehicle})
    return err
}

func (r *VehicleRepository) Delete(id string) error {
    objID, _ := primitive.ObjectIDFromHex(id)
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := r.Collection.DeleteOne(ctx, bson.M{"_id": objID})
    return err
}