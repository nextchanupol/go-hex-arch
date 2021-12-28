package repositoryports

import domainadapter "github.com/nextchanupol/go-fiber-server/internal/core/domain/adapters"

type CustomerRepository interface {
	GetAll() ([]domainadapter.Customer, error)
	GetByID(id int64) (*domainadapter.Customer, error)
}
