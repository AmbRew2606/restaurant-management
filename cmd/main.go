package main

import (
	"log"
	"net/http"
	"restaurant-management/config"
	"restaurant-management/internal/handlers"
	"restaurant-management/internal/services"

	"github.com/rs/cors"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Подключаемся к базе данных
	services.InitDB(cfg)
	defer services.CloseDB()

	// Настраиваем маршруты
	mux := http.NewServeMux()
	mux.HandleFunc("/menu", handlers.GetMenuHandler)

	// Настраиваем CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Запускаем сервер
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", c.Handler(mux)))
}
