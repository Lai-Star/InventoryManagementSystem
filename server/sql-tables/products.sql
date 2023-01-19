-- Create products table

CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    product_description TEXT NOT NULL,
    product_sku VARCHAR(255) NOT NULL,
    product_colour VARCHAR(255) NOT NULL,
    product_category VARCHAR(255) NOT NULL,
    product_brand VARCHAR(255) NOT NULL,
    total_quantity INT NOT NULL DEFAULT 0
);

CREATE TABLE sizes (
    size_id SERIAL PRIMARY KEY,
    size_name VARCHAR(255) NOT NULL,
    size_quantity INT NOT NULL
);

CREATE TABLE product_sizes (
    product_id INT REFERENCES products(product_id) ON DELETE CASCADE,
    size_id INT REFERENCES sizes(size_id) ON DELETE CASCADE,
    PRIMARY KEY (product_id, size_id)
);


CREATE TABLE orders (
    order_id INT PRIMARY KEY,
    customer_id INT REFERENCES customers(customer_id),
    product_id INT REFERENCES products(product_id),
    order_quantity INT NOT NULL,
    order_date DATE NOT NULL,
    order_status VARCHAR(255) NOT NULL
);

CREATE TABLE customers (
    customer_id INT PRIMARY KEY,
    customer_name VARCHAR(255) NOT NULL,
    customer_address VARCHAR(255) NOT NULL,
    customer_phone VARCHAR(255) NOT NULL,
    customer_email VARCHAR(255) NOT NULL
);

CREATE TABLE sales (
    sale_id INT PRIMARY KEY,
    product_id INT REFERENCES products(product_id),
    customer_id INT REFERENCES customers(customer_id),
    sale_quantity INT NOT NULL,
    sale_date DATE NOT NULL,
    product_cost DECIMAL(10,2) NOT NULL,
    sale_price DECIMAL(10,2) NOT NULL,
    profit DECIMAL(10,2) NOT NULL
);

CREATE TABLE inventory (
    inventory_id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES products(product_id),
    inventory_quantity INTEGER NOT NULL,
    inventory_date DATE NOT NULL,
    average_inventory DECIMAL(10,2) NOT NULL,
    cost_of_goods_sold DECIMAL(10,2) NOT NULL,
    inventory_turnover DECIMAL(10,2) NOT NULL
);

CREATE TABLE transactions (
    transaction_id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES products(product_id),
    transaction_quantity INTEGER NOT NULL,
    transaction_type VARCHAR(255) NOT NULL,
    transaction_date DATE NOT NULL
);

