// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package datalayer

import (
	"context"
)

const countOrders = `-- name: CountOrders :one
SELECT count(*) FROM orders
`

func (q *Queries) CountOrders(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countOrders)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (
  customer_id, items, status
) VALUES (
  $1, $2, $3
)
RETURNING order_id
`

type CreateOrderParams struct {
	CustomerID string
	Items      string
	Status     string
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (int32, error) {
	row := q.db.QueryRow(ctx, createOrder, arg.CustomerID, arg.Items, arg.Status)
	var order_id int32
	err := row.Scan(&order_id)
	return order_id, err
}

const listOrders = `-- name: ListOrders :many
SELECT order_id, customer_id, items, status, created_at FROM orders
ORDER BY created_at DESC
`

func (q *Queries) ListOrders(ctx context.Context) ([]Order, error) {
	rows, err := q.db.Query(ctx, listOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.OrderID,
			&i.CustomerID,
			&i.Items,
			&i.Status,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listOrdersByStatus = `-- name: ListOrdersByStatus :many
SELECT order_id, customer_id, items, status, created_at FROM orders
WHERE status LIKE $1
ORDER BY  created_at DESC
`

func (q *Queries) ListOrdersByStatus(ctx context.Context, status string) ([]Order, error) {
	rows, err := q.db.Query(ctx, listOrdersByStatus, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.OrderID,
			&i.CustomerID,
			&i.Items,
			&i.Status,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrderStatus = `-- name: UpdateOrderStatus :one
UPDATE orders
set status = $2
WHERE order_id = $1
RETURNING order_id, customer_id, items, status, created_at
`

type UpdateOrderStatusParams struct {
	OrderID int32
	Status  string
}

func (q *Queries) UpdateOrderStatus(ctx context.Context, arg UpdateOrderStatusParams) (Order, error) {
	row := q.db.QueryRow(ctx, updateOrderStatus, arg.OrderID, arg.Status)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.CustomerID,
		&i.Items,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}
