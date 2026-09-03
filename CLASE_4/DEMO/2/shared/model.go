package shared

type Pedido struct {
	PedidoID string  `json:"pedido_id"`
	Monto    float64 `json:"monto"`
	Cliente  string  `json:"cliente"`
}
