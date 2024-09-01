package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/joseasousa/desafio_cloud_run/internal/domain"
)

type WeatherRepository interface {
	GetWeatherByLocation(location string) (*domain.Weather, error)
}

type weatherRepository struct {
	client *http.Client
	url    string
	apiKey string
}

func NewWeatherRepository(client *http.Client, url, apiKey string) WeatherRepository {
	return &weatherRepository{
		client: client,
		url:    url,
		apiKey: apiKey,
	}
}

func (r *weatherRepository) GetWeatherByLocation(location string) (*domain.Weather, error) {
	escapedLocation := url.QueryEscape(location)
	url := fmt.Sprintf("%s?key=%s&q=%s", r.url, r.apiKey, escapedLocation)
	resp, err := r.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var weatherResponse domain.WeatherResponse
		if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
			return nil, err
		}
		response := domain.Weather{
			TempCelsius:    weatherResponse.Current.TempCelsius,
			TempFahrenheit: weatherResponse.Current.TempFahrenheit,
			TempKelvin:     weatherResponse.Current.TempCelsius + 273.15,
		}

		return &response, nil
	}

	return nil, fmt.Errorf("could not fetch weather data")
}
