package dto

// TemperaturaActual representa el objeto current que envía Open-Meteo.
type TemperaturaActual struct {
	TemperaturaC float64 `json:"temperature_2m"`
}

// RespuestaOpenMeteo representa los datos externos que necesitamos leer.
type RespuestaOpenMeteo struct {
	Current TemperaturaActual `json:"current"`
}

// RespuestaClima representa el JSON que devuelve nuestra API.
type RespuestaClima struct {
	Ciudad       string  `json:"ciudad"`
	TemperaturaC float64 `json:"temperatura_c"`
	Servidor     string  `json:"servidor"`
	Mensaje      string  `json:"mensaje"`
}
