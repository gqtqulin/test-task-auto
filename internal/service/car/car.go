package carserv

import (
	"github.com/gqtqulin/test-task-auto/internal/domain"
	"github.com/gqtqulin/test-task-auto/internal/storage"
)

type CarService struct {
	storage *storage.Storage
}

func NewCarService(storage *storage.Storage) *CarService {
	return &CarService{
		storage: storage,
	}
}

func (c *CarService) Create(car *domain.Car) (uint, error) {
	return c.storage.Car.Create(car)
}

func (c *CarService) Get(id uint) (*domain.Car, error) {
	return c.storage.Car.Get(id)
}

func (c *CarService) GetAll() ([]domain.Car, error) {
	return c.storage.Car.GetAll()
}

func (c *CarService) Delete(id uint) error {
	return c.storage.Car.Delete(id)
}
