package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gqtqulin/test-task-auto/internal/logger"
	"github.com/gqtqulin/test-task-auto/internal/service"
)

type Handler struct {
	carService service.Car
	log        *logger.Logger
}

func NewHandler(carService service.Car, log *logger.Logger) *Handler {
	return &Handler{
		carService: carService,
		log:        log,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/ping", h.Ping)

	cars := router.Group("/cars")
	{
		cars.GET("/", h.GetAllCars)
		cars.GET("/:id", h.GetCar)
		cars.POST("/add", h.AddCar)
		cars.DELETE("/:id", h.DeleteCar)
	}

	return router
}
