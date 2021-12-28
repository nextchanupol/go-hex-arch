package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/nextchanupol/go-fiber-server/internal/core/domain"
	domainadapter "github.com/nextchanupol/go-fiber-server/internal/core/domain/adapters"
	repositoryports "github.com/nextchanupol/go-fiber-server/internal/core/repositories/ports"
)

type customerRepositoryMock struct {
	customers []domainadapter.Customer
}

func NewCustomerRepositoryMock() repositoryports.CustomerRepository {
	customers := []domainadapter.Customer{
		{
			Customer: domain.Customer{
				CustomerID:  1,
				FirstName:   "John",
				LastName:    "Doe",
				DateOfBirth: "1986-01-01",
				Email:       sql.NullString{"test@test.com", true},
				IsActive:    true,
				CreatedDate: "2020-01-01",
				LastUpdate:  &time.Time{},
				NoOfActive:  sql.NullInt16{1, true},
			},
		},
		{
			Customer: domain.Customer{
				CustomerID:  2,
				FirstName:   "John2",
				LastName:    "Doe2",
				DateOfBirth: "1986-01-01",
				Email:       sql.NullString{"test2@test2.com", true},
				IsActive:    true,
				CreatedDate: "2020-01-02",
				LastUpdate:  &time.Time{},
				NoOfActive:  sql.NullInt16{5, true},
			},
		},
	}
	return repositoryports.CustomerRepository(customerRepositoryMock{customers: customers})
}

func (r customerRepositoryMock) GetAll() ([]domainadapter.Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetByID(id int64) (*domainadapter.Customer, error) {
	for _, v := range r.customers {
		if v.Customer.CustomerID == id {
			return &v, nil
		}
	}

	return nil, errors.New("customer not found")
}
