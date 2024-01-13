package main

import (
	"car-services-api.totote05.ar/api/handlers"
	"car-services-api.totote05.ar/repositories"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	vehicleAdapter := repositories.NewVehicle()
	vehicleHandler := handlers.NewVehicleHandler(router, vehicleAdapter)
	vehicleHandler.Init()

	router.Run()
}
