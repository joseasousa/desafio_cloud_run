package controller

import (
	"encoding/json"
	"net/http"

	"github.com/joseasousa/desafio_cloud_run/service"
)

type WeatherController struct {
	weatherService service.WeatherService
	zipCodeService service.ZipCodeService
}

func NewWeatherController(weatherService service.WeatherService, zipCodeService service.ZipCodeService) *WeatherController {
	return &WeatherController{
		weatherService: weatherService,
		zipCodeService: zipCodeService,
	}
}

func (h *WeatherController) GetWeather(w http.ResponseWriter, r *http.Request) {
	zipCode := r.PathValue("zipcode")

	iz := "invalid zipcode\n"
	cz := "can not find zipcode\n"
	if len(zipCode) != 8 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(iz))

		return
	}

	location, err := h.zipCodeService.GetLocationByZipCode(zipCode)
	if err != nil {
		if err.Error() == "can not find zipcode" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(cz))
		} else {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(iz))
		}
		return
	}

	weather, err := h.weatherService.GetWeatherByLocation(location.Localidade)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(cz))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(weather)

}
