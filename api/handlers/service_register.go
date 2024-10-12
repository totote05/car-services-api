package handlers

import (
	"net/http"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
	"car-services-api.totote05.ar/domain/usecases"
)

type (
	ServiceRegisterHandler crudHandler
	serviceRegisterHandler struct {
		serviceRegisterAdapter adapters.ServiceRegister
		vehicleAdapter         adapters.Vehicle
		serviceAdapter         adapters.Service
		kmAdapter              adapters.Km
	}
	createServiceRegisterPayload struct {
		ServiceID entities.ServiceID `json:"service_id"`
		Km        entities.Km        `json:"km"`
	}
)

func NewServiceRegisterHandler(
	serviceRegisterAdapter adapters.ServiceRegister,
	vehicleAdapter adapters.Vehicle,
	serviceAdapter adapters.Service,
	kmAdapter adapters.Km,
) ServiceRegisterHandler {
	return &serviceRegisterHandler{
		serviceRegisterAdapter: serviceRegisterAdapter,
		vehicleAdapter:         vehicleAdapter,
		serviceAdapter:         serviceAdapter,
		kmAdapter:              kmAdapter,
	}
}

func (sr *serviceRegisterHandler) List(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	id, err := GetID(r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	vehicleID := entities.VehicleID(id)

	usecase := usecases.NewGetRegisteredServices(
		sr.serviceRegisterAdapter,
		sr.vehicleAdapter,
	)

	result, err := usecase.Execute(c, vehicleID)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	JSON(w, http.StatusOK, result)
}

func (sr *serviceRegisterHandler) Create(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	id, err := GetID(r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	vehicleID := entities.VehicleID(id)

	payload, err := GetBody[createServiceRegisterPayload](r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	usecase := usecases.NewRegisterService(
		sr.serviceRegisterAdapter,
		sr.vehicleAdapter,
		sr.serviceAdapter,
		sr.kmAdapter,
	)

	result, err := usecase.Execute(c, vehicleID, payload.ServiceID, payload.Km)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	JSON(w, http.StatusCreated, result)
}

func (sr *serviceRegisterHandler) Get(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	id, err := GetID(r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	vehicleID := entities.VehicleID(id)
	serviceRegisterID, err := GetParam(r, "register_id")
	if err != nil {
		BadRequest(w, err)
		return
	}

	usecase := usecases.NewGetServiceRegister(
		sr.serviceRegisterAdapter,
		sr.vehicleAdapter,
	)

	result, err := usecase.Execute(c, vehicleID, entities.ServiceRegisterID(serviceRegisterID))
	if err != nil {
		InternalServerError(w, err)
		return
	}

	JSON(w, http.StatusOK, result)
}

func (sr *serviceRegisterHandler) Delete(w http.ResponseWriter, r *http.Request) {
	Unimplemented(w)
}

func (sr *serviceRegisterHandler) Update(w http.ResponseWriter, r *http.Request) {
	Unimplemented(w)
}
