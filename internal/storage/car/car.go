package carstore

import (
	"fmt"
	"github.com/gqtqulin/test-task-auto/internal/domain"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
)

const (
	carTable = "car"
)

type CarStorage struct {
	conn *pgx.Conn
}

func NewCarStorage(conn *pgx.Conn) *CarStorage {
	return &CarStorage{
		conn: conn,
	}
}

func (c *CarStorage) Create(car *domain.Car) (uint, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s (mark, model, owner_count, price, currency, options)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING car_id;
	`, carTable)
	var id uint

	err := c.conn.QueryRow(query,
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

func (c *CarStorage) Get(id uint) (*domain.Car, error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE car_id = $1;`, carTable)
	var options pgtype.VarcharArray

	car := &domain.Car{}
	err := c.conn.QueryRow(query,
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
		return nil, err
	}

	if options.Status == pgtype.Present {
		car.Options = make([]string, len(options.Elements))
		for i, el := range options.Elements {
			car.Options[i] = el.String
		}
	}

	return car, nil
}

func (c *CarStorage) GetAll() ([]domain.Car, error) {
	query := fmt.Sprintf(`SELECT * FROM %s;`, carTable)
	rows, err := c.conn.Query(query)
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

func (c *CarStorage) Delete(id uint) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE car_id = $1;`, carTable)

	tag, err := c.conn.Exec(query, id)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("car not found: %d", id)
	}

	return nil
}
