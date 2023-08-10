CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE NOT NULL,
    order_name TEXT NOT NULL,
    order_description TEXT NOT NULL,
    order_tracking_number INT UNIQUE NOT NULL
);