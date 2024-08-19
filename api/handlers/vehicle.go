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
	ErrMissingVehicleID = errors.New("error missing vehicle id")
)

type (
	Vehicle struct {
		router         *gin.Engine
		vehicleAdapter adapters.Vehicle
	}
)

func AddVehicleHandler(router *gin.Engine, vehicleAdapter adapters.Vehicle) {
	h := Vehicle{
		vehicleAdapter: vehicleAdapter,
	}

	group := h.router.Group("/vehicle")
	group.GET("/", h.List)
	group.POST("/", h.Create)
	group.GET("/:id", h.Get)
	group.PUT("/:id", h.Update)
	group.DELETE("/:id", h.Remove)
}

func (h Vehicle) getVehicleParam(c *gin.Context) (entities.VehicleID, error) {
	var (
		vehicleID entities.VehicleID
		err       error
	)

	id, ok := c.Params.Get("id")
	vehicleID = entities.VehicleID(id)
	if !ok {
		err = ErrMissingVehicleID
	}

	return vehicleID, err
}

func (h Vehicle) Create(c *gin.Context) {
	vehicle := entities.Vehicle{}
	if err := c.Bind(&vehicle); err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}

	usecase := usecases.NewCreateVehicle(h.vehicleAdapter)
	created, err := usecase.Execute(c, vehicle)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, created)
}

func (h Vehicle) Get(c *gin.Context) {
	vehicleID, err := h.getVehicleParam(c)
	if err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}

	usecase := usecases.NewGetVehicle(h.vehicleAdapter)
	vehicle, err := usecase.Execute(c, vehicleID)

	if errors.Is(err, adapters.ErrNotFound) {
		HandleError(c, http.StatusNotFound, err)
		return
	}

	if err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, vehicle)
}

func (h Vehicle) Update(c *gin.Context) {
	vehicleID, err := h.getVehicleParam(c)
	if err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}

	vehicle := entities.Vehicle{}
	if err := c.Bind(&vehicle); err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}

	vehicle.ID = vehicleID

	usecase := usecases.NewUpdateVehicle(h.vehicleAdapter)
	updatedVehicle, err := usecase.Execute(c, vehicle)

	if errors.Is(err, adapters.ErrNotFound) {
		HandleError(c, http.StatusNotFound, err)
		return
	}

	if err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, updatedVehicle)
}

func (h Vehicle) Remove(c *gin.Context) {
	vehicleID, err := h.getVehicleParam(c)
	if err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}

	usecase := usecases.NewDeleteVehicle(h.vehicleAdapter)
	if err := usecase.Execute(c, vehicleID); err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(200)
}

func (h Vehicle) List(c *gin.Context) {
	usecase := usecases.NewGetVehicles(h.vehicleAdapter)
	list, err := usecase.Execute(c)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, list)
}
