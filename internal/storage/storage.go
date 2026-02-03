package storage

import (
	"github.com/jackc/pgx"
	"test-task-auto/internal/domain"
	carstore "test-task-auto/internal/storage/car"
)

type Car interface {
	Create(car *domain.Car) (uint, error)
	Get(id uint) (*domain.Car, error)
	GetAll() ([]domain.Car, error)
	Delete(id uint) error
}

type Storage struct {
	conn *pgx.Conn
	Car  Car
}

func NewStorage(conn *pgx.Conn) *Storage {
	return &Storage{
		conn: conn,
		Car:  carstore.NewCarStorage(conn),
	}
}
