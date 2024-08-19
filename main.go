package main

import (
	"car-services-api.totote05.ar/api/handlers"
	"car-services-api.totote05.ar/repositories/local"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	vehicleAdapter := local.NewVehicle()
	handlers.AddVehicleHandler(router, vehicleAdapter)

	serviceAdapter := local.NewService()
	handlers.AddServiceHandler(router, serviceAdapter)

	router.Run()
}
