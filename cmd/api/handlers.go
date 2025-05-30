package main

import (
	"context"
	"fmt"
	"go-order-process/datalayer"
	"log"
	"net/http"
)

func (app *Config) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	var allOrders []datalayer.Order
	var err error
	isErr := false

	// check for status param
	ctx := r.Context()
	orderStatus, ok := ctx.Value(contextKey("validatedStatus")).(string)
	log.Printf("GET, context status: %s, ok: %t", orderStatus, ok)
	if len(orderStatus) > 0 && ok {
		// return all orders
		allOrders, err = app.Repo.ListOrdersByStatus(context.Background(), orderStatus)
		if err != nil {
			log.Printf("Error while fetching orders with status:%s. Error:%s\n",orderStatus, err)
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
	log.Println("Is valid Post request:", ok)
	if ok {
		fmt.Println("Valid POST Request:", validRequest)
	} else {
		fmt.Println("CTX value error")
	}
}