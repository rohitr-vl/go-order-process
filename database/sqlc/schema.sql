CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    customer_id VARCHAR (50) NOT NULL,
    items VARCHAR (255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL
);