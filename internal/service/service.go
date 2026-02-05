package service

import (
	"github.com/gqtqulin/test-task-auto/internal/domain"
	carserv "github.com/gqtqulin/test-task-auto/internal/service/car"
	"github.com/gqtqulin/test-task-auto/internal/storage"
)

type CarService interface {
	Create(*domain.Car) (uint, error)
	Get(id uint) (*domain.Car, error)
	GetAll() ([]domain.Car, error)
	Delete(id uint) error
}

type Service struct {
	storage *storage.Storage
	Car     CarService
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		storage: storage,
		Car:     carserv.NewCarService(storage),
	}
}
