package service

import (
	"github.com/joseasousa/desafio_cloud_run/internal/domain"
	"github.com/joseasousa/desafio_cloud_run/internal/infra/repository"
)

//go:generate mockery --name ZipCodeService --outpkg mock --output mock --filename zipcode.go --with-expecter=true

type ZipCodeService interface {
	GetLocationByZipCode(zipCode string) (*domain.Location, error)
}

type zipCodeService struct {
	repository repository.ZipCodeRepository
}

func NewZipCodeService(repo repository.ZipCodeRepository) ZipCodeService {
	return &zipCodeService{
		repository: repo,
	}
}

func (s *zipCodeService) GetLocationByZipCode(zipCode string) (*domain.Location, error) {
	return s.repository.GetLocationByZipCode(zipCode)
}
