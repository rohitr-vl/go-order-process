-- name: ListOrders :many
SELECT * FROM orders
ORDER BY created_at DESC;

-- name: ListOrdersByStatus :many
SELECT * FROM orders
WHERE status LIKE $1
ORDER BY  created_at DESC;

-- name: CreateOrder :one
INSERT INTO orders (
  customer_id, items, status
) VALUES (
  $1, $2, $3
)
RETURNING order_id;

-- name: UpdateOrderStatus :one
UPDATE orders
set status = $2
WHERE order_id = $1
RETURNING *;

-- name: CountOrders :one
SELECT count(*) FROM orders;