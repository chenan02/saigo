
CREATE TABLE customers(
customer_id SERIAL PRIMARY KEY
, email text UNIQUE
, address text
, ssn INTEGER UNIQUE
);

CREATE TABLE orders(
order_id SERIAL PRIMARY KEY
, product text
, quantity INTEGER
, customer_id INTEGER NOT NULL
, created_at timestamp without time zone DEFAULT timezone('utc'::text, now())
, CONSTRAINT orders_customer_id_fkey FOREIGN KEY (customer_id)
      REFERENCES customers (customer_id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
);
