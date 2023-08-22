package bad

import (
	"github.com/pkg/errors"
)

// errors
var (
	StrControllerFail = "controller: fail"
	StrProviderFail   = "provider: fail"
	StrServiceFail    = "service: fail"
	StrRepositoryFail = "repository: fail"

	ErrNotFound = errors.New("not found")
)

type Repository struct{}

func (r Repository) Find() error {
	err := ErrNotFound
	return errors.Wrap(err, StrRepositoryFail)
}

type Service struct {
	repository Repository
}

func (s Service) Execute() error {
	err := s.repository.Find()
	if err != nil {
		return errors.Wrap(err, StrServiceFail)
	}

	return nil
}

type Provider struct {
	service Service
}

func (p Provider) Execute() error {
	err := p.service.Execute()
	if err != nil {
		return errors.Wrap(err, StrProviderFail)
	}

	return nil
}

type Controller struct {
	provider Provider
}

func (c Controller) Execute() error {
	err := c.provider.Execute()
	if err != nil {
		return errors.Wrap(err, StrControllerFail)
	}

	return nil
}
