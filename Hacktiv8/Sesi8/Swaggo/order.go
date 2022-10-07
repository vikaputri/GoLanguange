package main

import (
	"encoding/json"
	"log"
	"net/http"
	"pkg/mod/github.com/gorilla/mux@v1.8.0"
	"strconv"
	"time"
)

type Order struct {
	OrderID      string    `json:"orderId" example:"1"`
	CustomerName string    `json:"customerName" example:"Vika"`
	OrderedAt    time.Time `json:"orderedAt" example:"2019-11-09T21:21:46+00:00"`
	Items        []Item    `json:"items`
}

type Item struct {
	ItemID      string `json:"itemId" example:"A1B2C3"`
	Description string `json:"description" example:"A random description"`
	Quantity    int    `json:"quantity" example:"1swag"`
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func createOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	prevOrderID++
	order.OrderID = strconv.Itoa(prevOrderID)
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/orders", createOrders).Methods("POST")
	router.HandleFunc("/orders/{orderId}", getOrders).Methods("Get")

	log.Fatal(http.ListenAndServe(":8080", router))
}
