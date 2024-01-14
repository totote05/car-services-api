package main

import (
	"car-services-api.totote05.ar/api/handlers"
	"car-services-api.totote05.ar/repositories/local"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	vehicleAdapter := local.NewVehicle()
	vehicleHandler := handlers.NewVehicleHandler(router, vehicleAdapter)
	vehicleHandler.Init()

	serviceAdapter := local.NewService()
	handlers.AddServiceHandler(router, serviceAdapter)

	router.Run()
}
