package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joseasousa/desafio_cloud_run/internal/application/config"
	"github.com/joseasousa/desafio_cloud_run/internal/infra/client"
	"github.com/joseasousa/desafio_cloud_run/internal/infra/entrypoint/controller"
	"github.com/joseasousa/desafio_cloud_run/internal/infra/repository"
	"github.com/joseasousa/desafio_cloud_run/service"
)

func main() {
	conf := config.NewConfig()
	httpClient := client.NewHTTPClient()

	mux := http.NewServeMux()

	zipCodeRepo := repository.NewZipCodeRepository(httpClient, conf.ViaCepURL)
	weatherRepo := repository.NewWeatherRepository(httpClient, conf.WeatherAPIURL, conf.WeatherAPIKey)

	zipCodeService := service.NewZipCodeService(zipCodeRepo)
	weatherService := service.NewWeatherService(weatherRepo)

	weatherController := controller.NewWeatherController(weatherService, zipCodeService)

	mux.HandleFunc("/clima/{zipcode}", weatherController.GetWeather)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	fmt.Println("rodando 8080")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
