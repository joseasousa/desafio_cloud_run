package domain

type WeatherResponse struct {
	Current Weather `json:"current"`
}

type Weather struct {
	TempCelsius    float64 `json:"temp_C"`
	TempFahrenheit float64 `json:"temp_F"`
	TempKelvin     float64 `json:"temp_K"`
}
