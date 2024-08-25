package handlers

import (
	"net/http"

	"car-services-api.totote05.ar/domain/adapters"
)

func NewServer(
	vehicleAdapter adapters.Vehicle,
	serviceAdapter adapters.Service,
	kmAdapter adapters.Km,
) *http.Server {
	mux := http.NewServeMux()

	serviceHandler := NewServiceHandler(serviceAdapter)
	mux.HandleFunc("GET /service", serviceHandler.List)
	mux.HandleFunc("POST /service", serviceHandler.Create)
	mux.HandleFunc("GET /service/{id}", serviceHandler.Get)
	mux.HandleFunc("PUT /service/{id}", serviceHandler.Update)
	mux.HandleFunc("DELETE /service/{id}", serviceHandler.Delete)

	vehicleHandler := NewVehicleHandler(vehicleAdapter)
	mux.HandleFunc("GET /vehicle", vehicleHandler.List)
	mux.HandleFunc("POST /vehicle", vehicleHandler.Create)
	mux.HandleFunc("GET /vehicle/{id}", vehicleHandler.Get)
	mux.HandleFunc("PUT /vehicle/{id}", vehicleHandler.Update)
	mux.HandleFunc("DELETE /vehicle/{id}", vehicleHandler.Delete)

	kmHandler := NewKmHandler(kmAdapter, vehicleAdapter)
	mux.HandleFunc("GET /vehicle/{id}/km", kmHandler.List)
	mux.HandleFunc("POST /vehicle/{id}/km", kmHandler.Create)
	mux.HandleFunc("GET /vehicle/{id}/km/{km_id}", kmHandler.Get)
	mux.HandleFunc("PUT /vehicle/{id}/km/{km_id}", kmHandler.Update)
	mux.HandleFunc("DELETE /vehicle/{id}/km/{km_id}", kmHandler.Delete)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
