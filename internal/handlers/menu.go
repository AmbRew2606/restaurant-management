package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"restaurant-management/internal/services"
)

func GetMenuItems(w http.ResponseWriter, r *http.Request) {
	items, err := services.FetchMenuItems()
	if err != nil {
		http.Error(w, "Failed to fetch menu items", http.StatusInternalServerError)
		return
	}

	// Парсинг поля Nutrition для каждого элемента
	for i := range items {
		items[i].ParseNutrition(items[i].Nutrition.(string))
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(items); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// package handlers

// import (
// 	"context"
// 	"encoding/json"
// 	"net/http"
// 	"restaurant-management/internal/services"
// )

// type MenuItem struct {
// 	ID          int     `json:"id"`
// 	Name        string  `json:"name"`
// 	Description string  `json:"description"`
// 	Nutrition   string  `json:"nutrition"`
// 	Price       float64 `json:"price"`
// }

// func GetMenuItems(w http.ResponseWriter, r *http.Request) {
// 	rows, err := services.DB.Query(context.Background(), "SELECT id, name, description, nutrition, price FROM menu_items")
// 	if err != nil {
// 		http.Error(w, "Failed to fetch menu items", http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	var menuItems []MenuItem

// 	for rows.Next() {
// 		var item MenuItem
// 		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Nutrition, &item.Price)
// 		if err != nil {
// 			http.Error(w, "Error scanning menu item", http.StatusInternalServerError)
// 			return
// 		}
// 		menuItems = append(menuItems, item)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(menuItems)
// }
