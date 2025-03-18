package main

import (
	"buffer-api-project/internal/buffer"
	"buffer-api-project/internal/worker"
	"buffer-api-project/pkg/model"
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// простой буфер в памяти
	factBuffer := buffer.NewMemoryBuffer(1000)
	// или Kafka-буфер
	// factBuffer := buffer.NewKafkaBuffer([]string{"localhost:9092"}, "facts_topic", "worker_group")

	wg.Add(1)
	go worker.Run(&wg, factBuffer, "48ab34464a5573519725deb5865cc74c")

	// эмуляция поступления данных
	for i := 0; i < 10; i++ {
		err := factBuffer.Push(
			model.Fact{
				PeriodStart:         "2025-03-15",
				PeriodEnd:           "2025-03-20",
				PeriodKey:           "month",
				IndicatorToMoId:     227373,
				IndicatorToMoFactId: 0,
				Value:               i + 1,
				FactTime:            "2025-03-18",
				IsPlan:              0,
				AuthUserId:          40,
				Comment:             fmt.Sprintf("buffer Nikitin #%d", i+1),
			})
		if err != nil {
			return
		}
	}

	err := factBuffer.Close() // это закроет канал и цикл завершится
	if err != nil {
		log.Printf("Ошибка закрытия буфера: %v", err)
		return
	}
	wg.Wait() // ждём завершения worker’ов
}
