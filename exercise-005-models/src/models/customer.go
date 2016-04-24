package models

import (
	"github.com/jmoiron/sqlx"
)

// Customer ...
type Customer struct {
	ID      int
	Email   string
	Address string
	SSN     int
	Orders  []*Order
}

// Refresh ...
func (c *Customer) Refresh(db *sqlx.DB) error {
	return nil
}

// NewCustomer ...
func NewCustomer(db *sqlx.DB, email string, address string, ssn int) (*Customer, error) {
	return nil, nil
}

// DeleteCustomer ...
func DeleteCustomer(db *sqlx.DB, id int) error {
	return nil
}

// UpdateCustomer ...
func UpdateCustomer(db *sqlx.DB, u *Customer) error {
	return nil
}

// FindCustomerByEmail ...
func FindCustomerByEmail(db *sqlx.DB, email string) (*Customer, error) {
	return nil, nil
}

// FindCustomerByID ...
func FindCustomerByID(db *sqlx.DB, id int) (*Customer, error) {
	return nil, nil
}

// AllCustomers ...
func AllCustomers(db *sqlx.DB) ([]*Customer, error) {
	return nil, nil
}
