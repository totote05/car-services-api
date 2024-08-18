package handlers

import (
	"errors"
	"net/http"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
	"car-services-api.totote05.ar/domain/usecases"
	"github.com/gin-gonic/gin"
)

var (
	ErrMissinServiceID = errors.New("error missing service id")
)

type (
	serviceHandler struct {
		serviceAdapter adapters.Service
	}
)

func AddServiceHandler(router *gin.Engine, serviceAdapter adapters.Service) {
	handlers := serviceHandler{
		serviceAdapter: serviceAdapter,
	}
	serviceRouter := router.Group("/service")
	serviceRouter.GET("/", handlers.List)
	serviceRouter.POST("/", handlers.Create)
	serviceRouter.GET("/:id", handlers.Get)
}

func (h *serviceHandler) List(c *gin.Context) {
	usecase := usecases.NewGetServices(h.serviceAdapter)

	list, err := usecase.Execute(c)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *serviceHandler) Create(c *gin.Context) {
	service := &entities.Service{}
	if err := c.Bind(service); err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}

	if err := service.Validate(); err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}

	usecase := usecases.NewCreateService(h.serviceAdapter)

	result, err := usecase.Execute(c, *service)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *serviceHandler) Get(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		HandleError(c, http.StatusBadRequest, ErrMissinServiceID)
	}

	serviceID := entities.ServiceID(id)

	usecase := usecases.NewGetService(h.serviceAdapter)
	result, err := usecase.Execute(c, serviceID)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, result)
}
