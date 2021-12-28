package repository

import (
	"github.com/jmoiron/sqlx"

	domainadapter "github.com/nextchanupol/go-fiber-server/internal/core/domain/adapters"
	repositoryports "github.com/nextchanupol/go-fiber-server/internal/core/repositories/ports"
)

type customerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) repositoryports.CustomerRepository {
	return repositoryports.CustomerRepository(customerRepositoryDB{db})
}

func (r customerRepositoryDB) GetAll() ([]domainadapter.Customer, error) {
	customers := []domainadapter.Customer{}
	query := `SELECT customer_id, first_name, last_name, date_of_birth, 
	email, is_active, created_date, last_update, no_of_active FROM customer`

	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r customerRepositoryDB) GetByID(id int64) (*domainadapter.Customer, error) {
	customer := domainadapter.Customer{}
	query := `SELECT customer_id, first_name, last_name, date_of_birth, 
	email, is_active, created_date, last_update, 
	no_of_active FROM customer where customer_id = $1`

	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
