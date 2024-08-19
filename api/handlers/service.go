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
	group := router.Group("/service")
	group.GET("/", handlers.List)
	group.POST("/", handlers.Create)
	group.GET("/:id", handlers.Get)
	group.PUT("/:id", handlers.Update)
	group.DELETE("/:id", handlers.Delete)
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
	id, err := h.getServiceParam(c)
	if err != nil {
		HandleError(c, http.StatusBadRequest, err)
	}

	serviceID := entities.ServiceID(id)

	usecase := usecases.NewGetService(h.serviceAdapter)
	result, err := usecase.Execute(c, serviceID)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, result)
}

func (h *serviceHandler) Update(c *gin.Context) {
	service := &entities.Service{}
	if err := c.Bind(service); err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}

	serviceID, err := h.getServiceParam(c)
	if err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}

	if serviceID != service.ID {
		HandleError(c, http.StatusBadRequest, errors.New("service id does not match"))
		return
	}

	usecase := usecases.NewUpdateService(h.serviceAdapter)
	result, err := usecase.Execute(c, *service)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *serviceHandler) Delete(c *gin.Context) {
	serviceID, err := h.getServiceParam(c)
	if err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}

	usecase := usecases.NewDeleteService(h.serviceAdapter)
	err = usecase.Execute(c, serviceID)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *serviceHandler) getServiceParam(c *gin.Context) (entities.ServiceID, error) {
	var (
		serviceID entities.ServiceID
		err       error
	)

	id, ok := c.Params.Get("id")
	serviceID = entities.ServiceID(id)
	if !ok {
		err = ErrMissinServiceID
	}

	return serviceID, err
}
