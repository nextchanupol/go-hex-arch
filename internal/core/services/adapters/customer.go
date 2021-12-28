package custsrv

import (
	"database/sql"
	"errors"
	"log"

	"github.com/nextchanupol/go-fiber-server/internal/core/dto"
	repositoryports "github.com/nextchanupol/go-fiber-server/internal/core/repositories/ports"
	serviceport "github.com/nextchanupol/go-fiber-server/internal/core/services/ports"
)

type customerService struct {
	custRepo repositoryports.CustomerRepository
}

func NewCustomerService(custRepo repositoryports.CustomerRepository) serviceport.CustomerService {
	return serviceport.CustomerService(customerService{custRepo: custRepo})
}

func (r customerService) GetCustomers() ([]dto.CustomerResponse, error) {
	customers, err := r.custRepo.GetAll()

	if err != nil {
		log.Println(err)
		return nil, err
	}
	custResponses := []dto.CustomerResponse{}
	for _, v := range customers {
		customer := dto.CustomerResponse{
			CustomerID:  v.CustomerID,
			FirstName:   v.FirstName,
			LastName:    v.LastName,
			DateOfBirth: v.DateOfBirth,
			Email:       v.Email.String,
			IsActive:    v.IsActive,
			NoOfActive:  int(v.NoOfActive.Int16),
		}
		custResponses = append(custResponses, customer)
	}
	return custResponses, nil
}

func (r customerService) GetCustomerByID(id int64) (*dto.CustomerResponse, error) {
	customer, err := r.custRepo.GetByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}
		log.Println(err)
		return nil, err
	}
	custResponse := dto.CustomerResponse{
		CustomerID:  customer.CustomerID,
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		DateOfBirth: customer.DateOfBirth,
		Email:       customer.Email.String,
		IsActive:    customer.IsActive,
		NoOfActive:  int(customer.NoOfActive.Int16),
	}
	return &custResponse, nil
}
