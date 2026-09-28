package dto

type TemperaturaActual struct {
	TemperaturaC float64 `json:"temperature_2m"`
}

type RespuestaOpenMeteo struct {
	Current TemperaturaActual `json:"current"`
}

type RespuestaClima struct {
	Ciudad       string  `json:"ciudad"`
	TemperaturaC float64 `json:"temperatura_c"`
	Servidor     string  `json:"servidor"`
	Mensaje      string  `json:"mensaje"`
}
