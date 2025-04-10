package config

import (
    "context"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*mongo.Client, *mongo.Collection) {
    uri := os.Getenv("MONGO_URI")
    if uri == "" {
        log.Fatal("❌ MONGO_URI no está definido en el archivo .env")
    }

    clientOptions := options.Client().ApplyURI(uri)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatalf("❌ Error conectando a MongoDB: %v", err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatalf("❌ Error al hacer ping a MongoDB: %v", err)
    }

    dbName := os.Getenv("MONGO_DB")
    collectionName := os.Getenv("MONGO_COLLECTION")

    if dbName == "" || collectionName == "" {
        log.Fatal("❌ MONGO_DB o MONGO_COLLECTION no están definidos en .env")
    }

    collection := client.Database(dbName).Collection(collectionName)

    return client, collection
}