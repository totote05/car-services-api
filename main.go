package main

import (
	"log"

	"car-services-api.totote05.ar/api/handlers"
	"car-services-api.totote05.ar/repositories/local"
)

func main() {

	vehicleAdapter := local.NewVehicle()
	serviceAdapter := local.NewService()
	kmAdapter := local.NewKm()
	serviceRegister := local.NewServiceRegister()

	server := handlers.NewServer(
		vehicleAdapter,
		serviceAdapter,
		kmAdapter,
		serviceRegister,
	)

	log.Println("Server running on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
