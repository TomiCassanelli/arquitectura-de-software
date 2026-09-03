package shared

const (
	AMQPURI        = "amqp://user:pass@localhost:5672"
	QueueName      = "cola-pedidos-resiliente"
	DLQName        = "cola-pedidos-dlq"
	MaxConnections = 3
	BaseDelayMs    = 500
)
