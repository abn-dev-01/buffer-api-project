package buffer

//
//	Имплементация буфера в виде канала в памяти.
//
import (
	"buffer-api-project/pkg/model"
	"errors"
)

type MemoryBuffer struct {
	ch chan model.Fact
}

func NewMemoryBuffer(size int) *MemoryBuffer {
	return &MemoryBuffer{ch: make(chan model.Fact, size)}
}

// Push отправляет факт в буфер
func (m *MemoryBuffer) Push(fact model.Fact) error {
	select {
	case m.ch <- fact:
		return nil
	default:
		return errors.New("буфер переполнен")
	}
}

// Read возвращает канал для чтения фактов
func (m *MemoryBuffer) Read() <-chan model.Fact {
	return m.ch
}

// Close закрывает канал буфера
func (m *MemoryBuffer) Close() error {
	close(m.ch)
	return nil
}
