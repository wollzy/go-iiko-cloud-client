package go_iiko_cloud

import (
	"github.com/dimonrus/goreq"
)

// Структура для сервиса documentor
type Service struct {
	// Запрос
	goreq.HttpRequest
	// Конфиг
	config Config
}

// Конфиг клиента
type Config struct {
	// Хост
	Host string `yaml:"host"`
}

// Создание экземпляра сервиса
func ServiceNew(r goreq.HttpRequest, cnf Config) Service {
	return Service{
		HttpRequest: r,
		config:      cnf,
	}
}
