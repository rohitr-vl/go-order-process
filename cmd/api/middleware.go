package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type contextKey string

func ValidateRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		log.Println("Middleware ValidateRequest, method:", method)
		switch method {
		case "GET":
			log.Println("Received Get request")
			statusParam := ValidateGetRequest(w, r)
			ctxRequestKey := contextKey("validatedStatus")
			ctx := context.WithValue(r.Context(), ctxRequestKey, statusParam)
			next.ServeHTTP(w, r.WithContext(ctx))
		case "POST":
			log.Println("Received Post request")
			req := ValidatePostRequest(w, r)
			ctxRequestKey := contextKey("validatedRequest")
			ctx := context.WithValue(r.Context(), ctxRequestKey, req)
			next.ServeHTTP(w, r.WithContext(ctx))
		default:
			fmt.Fprintf(w, "Unsupported method: %s\n", method)
		}
	})
}

func ValidateGetRequest(w http.ResponseWriter, r *http.Request) string {
		statusParam := chi.URLParam(r, "status")
		log.Printf("GET status param: %s\n", statusParam)
		statusParam = SanitizeValue(statusParam)
		log.Printf("GET status param after sanitization: %s\n", statusParam)
		if len(statusParam) > 0 {
			err := json.NewDecoder(r.Body).Decode(&statusParam)
			if err != nil {
				payload := jsonResponse{
					Error:   true,
					Message: fmt.Sprintf("Invalid order status error: %s", err),
					Data:    nil,
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(400)
				out, _ := json.Marshal(payload)
				log.Println("GET error payload:", err)
				_, err = w.Write(out)
				if err != nil {
					fmt.Println("Error while return error message from middleware")
				}
				return ""
			}
			return statusParam
		}
		return ""
	}

func ValidatePostRequest(w http.ResponseWriter, r *http.Request) jsonRequest {
		var req jsonRequest
		var errMsg string

		if r.Header["Content-Type"][0] != "application/json" {
			errMsg = "App can accept POST requests in json format only"
			fmt.Println(errMsg)
		}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errMsg = fmt.Sprintf("Error: %s", err)
			fmt.Println(errMsg)
		}
		custID := SanitizeValue(req.CustomerId)

		if len(custID) <= 0 {
			errMsg = "Please provide valid Customer ID!"
			fmt.Println(errMsg)
		}

		if len(req.Items) <= 0 {
			errMsg = "Please provide Order items!"
			fmt.Println(errMsg)
		}

		if req.Total <= 0 {
			errMsg = "Please provide Order price!"
			fmt.Println(errMsg)
		}

		if errMsg != "" {
			payload := jsonResponse{
				Error:   true,
				Message: errMsg,
				Data:    nil,
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			out, _ := json.Marshal(payload)
			_, err = w.Write(out)
			if err != nil {
				fmt.Println("Error while return error message from middleware")
			}
			// return req
		}
		// log.Println("Validated POST request: ", req)
		return req
	}