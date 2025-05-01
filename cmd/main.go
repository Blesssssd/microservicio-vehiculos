package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/gorilla/mux"

	"github.com/Blesssssd/microservicio-golang/config"
	"github.com/Blesssssd/microservicio-golang/controllers"
	"github.com/Blesssssd/microservicio-golang/repository"
)

func main() {
	godotenv.Load()
	client, collection := config.ConnectMongo()
	defer client.Disconnect(nil)

	fmt.Println("✅ Conectado a MongoDB")

	repo := repository.NewVehicleRepository(collection)
	ctrl := controllers.NewVehicleController(repo)

	r := mux.NewRouter()

	// ✅ Ruta adicional para "Hola Mundo"
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hola Mundo")
	}).Methods("GET")

	r.HandleFunc("/vehicles", ctrl.CreateVehicle).Methods("POST")
	r.HandleFunc("/vehicles", ctrl.GetVehicles).Methods("GET")
	r.HandleFunc("/vehicles/{id}", ctrl.GetVehicleByID).Methods("GET")
	r.HandleFunc("/vehicles/{id}", ctrl.UpdateVehicle).Methods("PUT")
	r.HandleFunc("/vehicles/{id}", ctrl.DeleteVehicle).Methods("DELETE")

	fmt.Println("🚗 Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
