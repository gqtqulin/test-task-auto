package storage

import (
	"github.com/gqtqulin/test-task-auto/internal/domain"
	carstore "github.com/gqtqulin/test-task-auto/internal/storage/car"
	"github.com/jackc/pgx"
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
