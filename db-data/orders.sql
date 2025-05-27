-- create schema and user and grant permissions
CREATE SCHEMA dev_rr_vl;
create user dev_rr_user with encrypted password 'super369';
grant all privileges on database "go-order-process" to dev_rr_user;
alter user dev_rr_user set search_path=dev_rr_vl, “$user”, public;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA dev_rr_vl TO dev_rr_user;



SET default_tablespace = '';
SET default_table_access_method = heap;

--
-- Name: orders; Type: TABLE; Schema: dev_rr_vl; Owner: dev_rr_user
--

CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    customer_id VARCHAR (50) NOT NULL,
    items VARCHAR (255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL
);
