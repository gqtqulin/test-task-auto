package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gqtqulin/test-task-auto/internal/domain"
	"net/http"
	"strconv"
)

func (h *Handler) GetAllCars(c *gin.Context) {
	cars, err := h.service.Car.GetAll()
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, cars)
}

func (h *Handler) GetCar(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		errorResponse(c, http.StatusBadRequest, "id is required")
		return
	}

	numId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "id must be numeric")
		return
	}

	car, err := h.service.Car.Get(uint(numId))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, car)
}

func (h *Handler) AddCar(c *gin.Context) {
	var car domain.Car

	err := c.ShouldBindBodyWithJSON(&car)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "json format error")
		return
	}

	id, err := h.service.Car.Create(&car)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) DeleteCar(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		errorResponse(c, http.StatusBadRequest, "id is required")
		return
	}

	numId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "id must be numeric")
		return
	}

	err = h.service.Car.Delete(uint(numId))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
