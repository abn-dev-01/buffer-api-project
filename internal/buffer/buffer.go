package buffer

import "buffer-api-project/pkg/model"

// Buffer интерфейс абстрактного буфера.
// Имплементацией работы с буфером может быть работа с Kafka, RabbitMQ, Memory.
type Buffer interface {
	Push(fact model.Fact) error
	Read() <-chan model.Fact
	Close() error
}
