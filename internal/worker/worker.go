package worker

import (
	"buffer-api-project/internal/buffer"
	"buffer-api-project/internal/client"
	"fmt"
	"log"
	"sync"
)

/*
 	wg *sync.WaitGroup: позволяет главной программе дождаться завершения работы worker’а.
	buf buffer.Buffer: интерфейс, абстрагирующий источник данных (буфер, который может быть в памяти, Kafka, Redis и т.п.).
	token string: токен авторизации для внешнего API.
*/
// Run универсальный worker, принимает любой буфер
func Run(wg *sync.WaitGroup, buf buffer.Buffer, token string) {
	// Сообщаем WaitGroup’у, что когда функция завершится, нужно уменьшить счётчик активных горутин на 1.
	defer wg.Done()

	// Клиент заранее настраивается на работу с внешним API
	apiClient := client.NewAPIClient(token)

	// Основной цикл обработки данных
	for fact := range buf.Read() {
		// Отправка данных на внешний API
		resp, err := apiClient.R().
			// Данные из структуры fact преобразуются в формат x-www-form-urlencoded и отправляются методом POST
			SetFormData(
				map[string]string{
					"period_start":            fact.PeriodStart,
					"period_end":              fact.PeriodEnd,
					"period_key":              fact.PeriodKey,
					"indicator_to_mo_id":      fmt.Sprintf("%d", fact.IndicatorToMoId),
					"indicator_to_mo_fact_id": fmt.Sprintf("%d", fact.IndicatorToMoFactId),
					"value":                   fmt.Sprintf("%d", fact.Value),
					"fact_time":               fact.FactTime,
					"is_plan":                 fmt.Sprintf("%d", fact.IsPlan),
					"auth_user_id":            fmt.Sprintf("%d", fact.AuthUserId),
					"comment":                 fact.Comment,
				}).
			Post("/_api/facts/save_fact")

		// Чтобы корректно завершить worker, нужно закрыть буфер
		if err != nil || resp.IsError() {
			log.Printf("Ошибка отправки на API: %v, resp: %s", err, resp)
			continue
		}

		log.Printf("Факт отправлен успешно: %s", resp.String())
	}
}
