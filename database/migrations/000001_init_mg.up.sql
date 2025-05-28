--
-- Name: orders; Type: TABLE; Schema: dev_rr_vl; Owner: dev_rr_user
--

CREATE TABLE IF NOT EXISTS orders (
    order_id SERIAL PRIMARY KEY,
    customer_id VARCHAR (50) NOT NULL,
    items VARCHAR (255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL
);