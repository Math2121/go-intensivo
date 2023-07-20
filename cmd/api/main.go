package main

import (
	"encoding/json"
	"net/http"

	"github.com/Math2121/go-first-step/internal/entity"

)


func main() {
	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/order", Hello)

	// http.ListenAndServe(":8888", r)

	e := echo.New() 
	e.GET("/order", Hello)

	e.Logger.Fatal(e.Start(":8888"))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	order, _ := entity.NewOrder("124", 200, 20)

	order.CalculateFinalPrice()

	json.NewEncoder(w).Encode(order)
}

func Order(c echo.Context) error {
	order := entity.Order{
		ID:    "5",
		Price: 20,
		Tax:   5,
	}

	err := order.CalculateFinalPrice()
	if err != nil {
		return nil
	}
	return c.JSON(http.StatusOK, order)
}
