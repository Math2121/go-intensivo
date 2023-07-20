package database

import (
	"database/sql"

	"github.com/Math2121/go-first-step/internal/entity"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	_, err := r.DB.Exec("Insert INTO orders (id, price, tax, final_price) VALUES (?,?,?,?)", order.ID, order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.DB.QueryRow("SELECT count(*) FROM orders").Scan(&total)

	if err != nil {
		return 0, err
	}
	return total, nil
}
