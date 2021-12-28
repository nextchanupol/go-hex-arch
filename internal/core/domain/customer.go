package domain

import (
	"database/sql"
	"time"
)

type Customer struct {
	CustomerID  int64          `db:"customer_id"`
	FirstName   string         `db:"first_name"`
	LastName    string         `db:"last_name"`
	DateOfBirth string         `db:"date_of_birth"`
	Email       sql.NullString `db:"email"`
	CreatedDate string         `db:"created_date"`
	IsActive    bool           `db:"is_active"`
	LastUpdate  *time.Time     `db:"last_update"`
	NoOfActive  sql.NullInt16  `db:"no_of_active"`
}
