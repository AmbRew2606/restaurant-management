package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"restaurant-management/internal/services"
)

type MenuItem struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Nutrition   string  `json:"nutrition"`
	Price       float64 `json:"price"`
}

func GetMenuItems(w http.ResponseWriter, r *http.Request) {
	rows, err := services.DB.Query(context.Background(), "SELECT id, name, description, nutrition, price FROM menu_items")
	if err != nil {
		http.Error(w, "Failed to fetch menu items", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var menuItems []MenuItem

	for rows.Next() {
		var item MenuItem
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Nutrition, &item.Price)
		if err != nil {
			http.Error(w, "Error scanning menu item", http.StatusInternalServerError)
			return
		}
		menuItems = append(menuItems, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menuItems)
}

// GetMenuHandler обрабатывает запросы к маршруту /menu
func GetMenuHandler(w http.ResponseWriter, r *http.Request) {
	menu := []map[string]interface{}{
		{"id": 1, "name": "Caesar Salad", "description": "Classic Caesar salad with chicken", "price": 10.99},
		{"id": 2, "name": "Margherita Pizza", "description": "Traditional pizza with mozzarella", "price": 14.99},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menu)
}
