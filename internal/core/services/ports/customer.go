package serviceport

import "github.com/nextchanupol/go-fiber-server/internal/core/dto"

type CustomerService interface {
	GetCustomers() ([]dto.CustomerResponse, error)
	GetCustomerByID(id int64) (*dto.CustomerResponse, error)
}
