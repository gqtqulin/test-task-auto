package service

import (
	"github.com/gqtqulin/test-task-auto/internal/domain"
	"github.com/gqtqulin/test-task-auto/internal/storage"
)

type Car interface {
	Create(domain.Car) (int, error)
	Get(id int) (domain.Car, error)
	GetAll() ([]domain.Car, error)
	Delete(id int) error
}

type CarService struct {
	storage storage.Car
}

func NewCarService(storage storage.Car) Car {
	return &CarService{
		storage: storage,
	}
}

func (s *CarService) Create(car domain.Car) (int, error) {
	return s.storage.Create(car)
}

func (s *CarService) Get(id int) (domain.Car, error) {
	return s.storage.Get(id)
}

func (s *CarService) GetAll() ([]domain.Car, error) {
	return s.storage.GetAll()
}

func (s *CarService) Delete(id int) error {
	return s.storage.Delete(id)
}
