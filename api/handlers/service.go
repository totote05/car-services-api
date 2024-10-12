package handlers

import (
	"errors"
	"net/http"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
	"car-services-api.totote05.ar/domain/usecases"
)

type (
	ServiceHandler crudHandler

	serviceHandler struct {
		serviceAdapter adapters.Service
	}
)

func NewServiceHandler(serviceAdapter adapters.Service) ServiceHandler {
	return &serviceHandler{
		serviceAdapter: serviceAdapter,
	}
}

func (h *serviceHandler) List(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	usecase := usecases.NewGetServices(h.serviceAdapter)

	list, err := usecase.Execute(c)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	JSON(w, http.StatusOK, list)
}

func (h *serviceHandler) Create(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	service, err := GetBody[entities.Service](r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	if err := service.Validate(); err != nil {
		BadRequest(w, err)
		return
	}

	usecase := usecases.NewCreateService(h.serviceAdapter)

	result, err := usecase.Execute(c, service)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	JSON(w, http.StatusCreated, result)
}

func (h *serviceHandler) Get(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	id, err := GetID(r)
	if err != nil {
		BadRequest(w, err)
	}

	serviceID := entities.ServiceID(id)

	usecase := usecases.NewGetService(h.serviceAdapter)
	result, err := usecase.Execute(c, serviceID)
	if err != nil {
		InternalServerError(w, err)
	}

	JSON(w, http.StatusOK, result)
}

func (h *serviceHandler) Update(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	id, err := GetID(r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	service, err := GetBody[entities.Service](r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	if entities.ServiceID(id) != service.ID {
		BadRequest(w, errors.New("service id does not match"))
		return
	}

	usecase := usecases.NewUpdateService(h.serviceAdapter)
	result, err := usecase.Execute(c, service)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	JSON(w, http.StatusOK, result)
}

func (h *serviceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	id, err := GetID(r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	serviceID := entities.ServiceID(id)
	usecase := usecases.NewDeleteService(h.serviceAdapter)
	err = usecase.Execute(c, serviceID)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
