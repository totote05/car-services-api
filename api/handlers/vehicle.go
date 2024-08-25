package handlers

import (
	"errors"
	"net/http"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
	"car-services-api.totote05.ar/domain/usecases"
)

type (
	VehicleHandler interface {
		List(w http.ResponseWriter, r *http.Request)
		Create(w http.ResponseWriter, r *http.Request)
		Get(w http.ResponseWriter, r *http.Request)
		Update(w http.ResponseWriter, r *http.Request)
		Delete(w http.ResponseWriter, r *http.Request)
	}

	vehicleHandler struct {
		vehicleAdapter adapters.Vehicle
	}
)

func NewVehicleHandler(vehicleAdapter adapters.Vehicle) VehicleHandler {
	return &vehicleHandler{
		vehicleAdapter: vehicleAdapter,
	}
}

func (h vehicleHandler) List(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	usecase := usecases.NewGetVehicles(h.vehicleAdapter)
	list, err := usecase.Execute(c)
	if err != nil {
		InternalServerError(w, err)
		return
	}
	JSON(w, http.StatusCreated, list)
}

func (h vehicleHandler) Create(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	vehicle, err := GetBody[entities.Vehicle](r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	usecase := usecases.NewCreateVehicle(h.vehicleAdapter)
	created, err := usecase.Execute(c, vehicle)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	JSON(w, http.StatusCreated, created)
}

func (h vehicleHandler) Get(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	id, err := GetID(r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	vehicleID := entities.VehicleID(id)

	usecase := usecases.NewGetVehicle(h.vehicleAdapter)
	vehicle, err := usecase.Execute(c, vehicleID)

	if errors.Is(err, adapters.ErrNotFound) {
		NotFound(w, err)
		return
	}

	if err != nil {
		InternalServerError(w, err)
		return
	}

	JSON(w, http.StatusCreated, vehicle)
}

func (h vehicleHandler) Update(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	id, err := GetID(r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	vehicle, err := GetBody[entities.Vehicle](r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	if vehicle.ID != entities.VehicleID(id) {
		BadRequest(w, errors.New("vehicle id does not match"))
		return
	}

	usecase := usecases.NewUpdateVehicle(h.vehicleAdapter)
	updatedVehicle, err := usecase.Execute(c, vehicle)

	if errors.Is(err, adapters.ErrNotFound) {
		NotFound(w, err)
		return
	}

	if err != nil {
		InternalServerError(w, err)
		return
	}

	JSON(w, http.StatusCreated, updatedVehicle)
}

func (h vehicleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	id, err := GetID(r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	vehicleID := entities.VehicleID(id)
	usecase := usecases.NewDeleteVehicle(h.vehicleAdapter)
	if err := usecase.Execute(c, vehicleID); err != nil {
		InternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
