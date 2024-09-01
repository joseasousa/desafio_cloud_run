package service

import (
	"github.com/joseasousa/desafio_cloud_run/internal/domain"
	"github.com/joseasousa/desafio_cloud_run/internal/infra/repository"
)

//go:generate mockery --name WeatherService --outpkg mock --output mock --filename weather.go --with-expecter=true

type WeatherService interface {
	GetWeatherByLocation(location string) (*domain.Weather, error)
}

type weatherService struct {
	repository repository.WeatherRepository
}

func NewWeatherService(repo repository.WeatherRepository) WeatherService {
	return &weatherService{
		repository: repo,
	}
}

func (s *weatherService) GetWeatherByLocation(location string) (*domain.Weather, error) {
	return s.repository.GetWeatherByLocation(location)
}
