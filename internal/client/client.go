package client

import (
	"github.com/go-resty/resty/v2"
)

// NewAPIClient возвращает настроенный HTTP-клиент
func NewAPIClient(token string) *resty.Client {
	return resty.New().
		SetHostURL("https://development.kpi-drive.ru").
		SetAuthToken(token).
		SetHeader("Content-Type", "application/x-www-form-urlencoded")
}
