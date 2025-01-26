package main

import (
	"log"
	"net/http"
	"restaurant-management/config"
	"restaurant-management/internal/handlers"
	"restaurant-management/internal/services"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Подключаемся к базе данных
	services.InitDB(cfg)
	defer services.CloseDB()

	// Настраиваем маршруты
	http.HandleFunc("/menu", handlers.GetMenuItems)

	// Запускаем сервер
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
