package worker

import (
	"buffer-api-project/internal/buffer"
	"buffer-api-project/internal/client"
	"fmt"
	"log"
	"sync"
)

// Run универсальный worker, принимает любой буфер
func Run(wg *sync.WaitGroup, buf buffer.Buffer, token string) {
	defer wg.Done()

	apiClient := client.NewAPIClient(token)

	for fact := range buf.Read() {
		resp, err := apiClient.R().
			SetFormData(map[string]string{
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

		if err != nil || resp.IsError() {
			log.Printf("Ошибка отправки на API: %v, resp: %s", err, resp)
			continue
		}

		log.Printf("Факт отправлен успешно: %s", resp.String())
	}
}
