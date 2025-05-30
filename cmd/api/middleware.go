package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type contextKey string

func ValidateRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		switch method {
		case "GET":
			fmt.Fprintf(w, "This is a GET request")
		case "POST":
			ValidatePostRequest(next)
		default:
			fmt.Fprintf(w, "Unsupported method: %s", method)
		}
	})
}

func ValidatePostRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
			return
		}

		ctxRequestKey := contextKey("validatedRequest")
		ctx := context.WithValue(r.Context(), ctxRequestKey, req)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
