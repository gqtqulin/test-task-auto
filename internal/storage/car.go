package storage

import (
	"fmt"
	"github.com/gqtqulin/test-task-auto/internal/domain"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
)

type CarStorage struct {
	conn *pgx.Conn
}

func NewCarStorage(conn *pgx.Conn) *CarStorage {
	return &CarStorage{
		conn: conn,
	}
}

func (s *CarStorage) Create(car domain.Car) (int, error) {
	query := `
		INSERT INTO car (mark, model, owner_count, price, currency, options)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING car_id;
	`
	var id int

	err := s.conn.QueryRow(query,
		car.Mark,
		car.Model,
		car.OwnerCount,
		car.Price,
		string(car.Currency),
		car.Options,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *CarStorage) Get(id int) (domain.Car, error) {
	query := `SELECT * FROM car WHERE car_id = $1;`
	var options pgtype.VarcharArray

	car := domain.Car{}
	err := s.conn.QueryRow(query,
		id,
	).Scan(
		&car.CarId,
		&car.Mark,
		&car.Model,
		&car.OwnerCount,
		&car.Price,
		&car.Currency,
		&options,
	)
	if err != nil {
		return car, err
	}

	if options.Status == pgtype.Present {
		car.Options = make([]string, len(options.Elements))
		for i, el := range options.Elements {
			car.Options[i] = el.String
		}
	}

	return car, nil
}

func (s *CarStorage) GetAll() ([]domain.Car, error) {
	query := `SELECT * FROM car;`
	rows, err := s.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []domain.Car
	for rows.Next() {
		var car domain.Car
		var options pgtype.VarcharArray

		if err := rows.Scan(
			&car.CarId,
			&car.Mark,
			&car.Model,
			&car.OwnerCount,
			&car.Price,
			&car.Currency,
			&options,
		); err != nil {
			return nil, err
		}

		// TODO: вынести в отдельную функцию
		if options.Status == pgtype.Present {
			car.Options = make([]string, len(options.Elements))
			for i, el := range options.Elements {
				car.Options[i] = el.String
			}
		}
		cars = append(cars, car)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cars, nil
}

func (s *CarStorage) Delete(id int) error {
	query := `DELETE FROM car WHERE car_id = $1;`

	tag, err := s.conn.Exec(query, id)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("car not found: %d", id)
	}

	return nil
}
