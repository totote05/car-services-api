package handlers

import (
	"errors"
	"net/http"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
	"car-services-api.totote05.ar/domain/usecases"
)

type (
	KmHandler interface {
		List(w http.ResponseWriter, r *http.Request)
		Create(w http.ResponseWriter, r *http.Request)
		Get(w http.ResponseWriter, r *http.Request)
		Update(w http.ResponseWriter, r *http.Request)
		Delete(w http.ResponseWriter, r *http.Request)
	}

	kmHandler struct {
		kmAdapter      adapters.Km
		vehicleAdapter adapters.Vehicle
	}
)

func NewKmHandler(kmAdapter adapters.Km, vehicleAdapter adapters.Vehicle) KmHandler {
	return &kmHandler{
		kmAdapter:      kmAdapter,
		vehicleAdapter: vehicleAdapter,
	}
}

func (h kmHandler) List(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	id, err := GetID(r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	vehicleID := entities.VehicleID(id)
	usecase := usecases.NewGetKms(h.kmAdapter, h.vehicleAdapter)
	list, err := usecase.Execute(c, vehicleID)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	JSON(w, http.StatusOK, list)
}

func (h kmHandler) Create(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	id, err := GetID(r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	km, err := GetBody[entities.Km](r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	vehicleID := entities.VehicleID(id)
	usecase := usecases.NewRegisterKm(h.kmAdapter, h.vehicleAdapter)
	result, err := usecase.Execute(c, vehicleID, km)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	JSON(w, http.StatusCreated, result)
}

func (h kmHandler) Get(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	id, err := GetID(r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	kmid, err := GetParam(r, "km_id")
	if err != nil {
		BadRequest(w, err)
		return
	}

	vehicleID := entities.VehicleID(id)
	kmID := entities.KmID(kmid)
	usecase := usecases.NewGetKm(h.kmAdapter, h.vehicleAdapter)
	km, err := usecase.Execute(c, vehicleID, kmID)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	JSON(w, http.StatusOK, km)
}

func (h kmHandler) Update(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	id, err := GetID(r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	kmid, err := GetParam(r, "km_id")
	if err != nil {
		BadRequest(w, err)
		return
	}

	km, err := GetBody[entities.Km](r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	vehicleID := entities.VehicleID(id)
	kmID := entities.KmID(kmid)

	if km.ID != kmID {
		BadRequest(w, errors.New("km id does not match"))
		return
	}
	usecase := usecases.NewUpdateKm(h.kmAdapter, h.vehicleAdapter)
	result, err := usecase.Execute(c, vehicleID, kmID, km)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	JSON(w, http.StatusOK, result)
}

func (h kmHandler) Delete(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	id, err := GetID(r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	kmid, err := GetParam(r, "km_id")
	if err != nil {
		BadRequest(w, err)
		return
	}

	vehicleID := entities.VehicleID(id)
	kmID := entities.KmID(kmid)
	usecase := usecases.NewDeleteKm(h.kmAdapter, h.vehicleAdapter)
	err = usecase.Execute(c, vehicleID, kmID)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	JSON(w, http.StatusNoContent, nil)
}
