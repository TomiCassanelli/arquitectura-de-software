package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"actividad-balanceador-roto/config"
	"actividad-balanceador-roto/dto"
)

// OpenMeteoService consulta la API externa de Open-Meteo.
type OpenMeteoService struct{}

// ObtenerTemperatura devuelve la temperatura actual o un error si la consulta falla.
func (OpenMeteoService) ObtenerTemperatura() (*dto.TemperaturaActual, error) {
	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%v&longitude=%v&current=temperature_2m",
		config.Latitud,
		config.Longitud,
	)

	cliente := http.Client{Timeout: 5 * time.Second}
	response, err := cliente.Get(url)
	if err != nil {
		return nil, fmt.Errorf("no se pudo consultar Open-Meteo: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Open-Meteo respondió estado %d", response.StatusCode)
	}

	var datos dto.RespuestaOpenMeteo
	if err := json.NewDecoder(response.Body).Decode(&datos); err != nil {
		return nil, fmt.Errorf("no se pudo leer la respuesta de Open-Meteo: %w", err)
	}

	return &datos.Current, nil
}
