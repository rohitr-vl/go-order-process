package main

import (
	"context"
	"go-order-process/datalayer"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (app *Config) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	var allOrders []datalayer.Order
	var err error
	var isErr bool
	// check for status param
	statusParam := chi.URLParam(r, "status")
	if len(strings.TrimSpace(statusParam)) == 0 {
		// return all orders
		allOrders, err = app.Repo.ListOrdersByStatus(context.Background(), statusParam)
		if err != nil {
			log.Printf("Error while fetching orders with status:%s. Error:%s\n",statusParam, err)
			isErr=true
		}
	} else {
		// search for orders based on given order status
		allOrders, err = app.Repo.ListOrders(context.Background())
		if err != nil {
			log.Printf("Error while fetching orders:%s\n", err)
			isErr=true
		}	
	}
	payload := jsonResponse{
		Error:   isErr,
		Message: err.Error(),
		Data:    allOrders,
	}
	app.writeJson(w, http.StatusAccepted, payload)
}

func (app *Config) AcceptOrder(w http.ResponseWriter, r *http.Request) {

}