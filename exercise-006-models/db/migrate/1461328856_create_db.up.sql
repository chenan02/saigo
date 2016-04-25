CREATE TABLE customers(
customer_id SERIAL PRIMARY KEY
, email VARCHAR(254)
, address VARCHAR(1000)
, ssn INTEGER
);

CREATE TABLE orders(
order_id SERIAL PRIMARY KEY
, product VARCHAR(254)
, quantity INTEGER
, customer_id INTEGER REFERENCES customers ON DELETE CASCADE
, created_at TIMESTAMPTZ default now()
);
