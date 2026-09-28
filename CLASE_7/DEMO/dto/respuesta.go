package dto

// RespuestaTest es el JSON que devuelve GET /test.
type RespuestaTest struct {
	Servidor string `json:"servidor"`
	Mensaje  string `json:"mensaje"`
}
