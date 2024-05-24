package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

// All CRUD handlers of Order

// Creates an order
func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create an order")
}

// Lists all the orders
func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all orders")
}

// Returns an order by ID
func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an order by ID")
}

// Updates an order by ID
func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an order by ID")
}

// Deletes an order by ID
func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an order by ID")
}
