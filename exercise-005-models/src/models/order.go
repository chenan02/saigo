package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// Order ...
type Order struct {
	ID         int
	Product    string
	Quantity   int
	Time       time.Time
	CustomerID int
}

// NewOrder ...
func NewOrder(db *sqlx.DB, custID int, product string, quantity int) error {
	return nil
}

// Update Order ...
func UpdateOrder(db *sqlx.DB, o *Order) error {
	return nil
}

// DeleteOrder ...
func DeleteOrder(db *sqlx.DB, orderID int) error {
	return nil
}
