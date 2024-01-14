package handlers

import (
	"net/http"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/usecases"
	"github.com/gin-gonic/gin"
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
