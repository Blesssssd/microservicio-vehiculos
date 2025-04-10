package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vehicle struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Brand       string             `bson:"brand" json:"brand"`
    Model       string             `bson:"model" json:"model"`
    Year        int                `bson:"year" json:"year"`
    LicensePlate string            `bson:"license_plate" json:"license_plate"`
}