package main

import (
	"context"
	"fmt"
	"go-order-process/datalayer"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Config) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	var allOrders []datalayer.Order
	var err error
	isErr := false

	// check for status param
	statusParam := chi.URLParam(r, "status")
	statusParam = SanitizeValue(statusParam)
	// log.Printf("GET param status: %s\n", statusParam)
	if len(statusParam) > 0 {
		// return all orders
		allOrders, err = app.Repo.ListOrdersByStatus(context.Background(), statusParam)
		if err != nil {
			log.Printf("Error while fetching orders with status:%s. Error:%s\n",statusParam, err)
			isErr=true
		} else {
			fmt.Println("All Orders by status:", allOrders)
		}
	} else {
		// search for orders based on given order status
		allOrders, err = app.Repo.ListOrders(context.Background())
		if err != nil {
			log.Printf("Error while fetching orders:%s\n", err)
			isErr=true
		} else {
			fmt.Println("All Orders:", allOrders)
		}
	}
	payload := jsonResponse{
		Error:   isErr,
		Message: fmt.Sprintf("Error:%s", err),
		Data:    allOrders,
	}
	app.writeJson(w, http.StatusAccepted, payload)
}

func (app *Config) AcceptOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	validRequest, ok := ctx.Value(contextKey("validatedRequest")).(jsonRequest)
	if ok {
		fmt.Println("Valid POST Request:", validRequest)
	} else {
		fmt.Println("CTX value error")
	}
}