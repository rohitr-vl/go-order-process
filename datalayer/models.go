// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package datalayer

import (
	"time"
)

type Order struct {
	OrderID    int32
	CustomerID string
	Items      string
	Status     string
	CreatedAt  time.Time
}
