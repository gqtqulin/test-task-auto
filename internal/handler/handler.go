package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gqtqulin/test-task-auto/internal/domain"
	"log/slog"
)

type CarService interface {
	Create(domain.Car) (int, error)
	Get(id int) (domain.Car, error)
	GetAll() ([]domain.Car, error)
	Delete(id int) error
}

type Handler struct {
	carService CarService
	log        *slog.Logger
}

func NewHandler(carService CarService, log *slog.Logger) *Handler {
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
