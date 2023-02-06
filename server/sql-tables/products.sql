-- Create products table
/*
    Table Description: To store details of the product
*/

CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    product_description VARCHAR(255),
    product_sku VARCHAR(50) NOT NULL,
    product_colour VARCHAR(50) NOT NULL,
    product_category VARCHAR(20) NOT NULL,
    product_brand VARCHAR(50) NOT NULL,
    product_cost DECIMAL(10,2) NOT NULL,
    added_date TIMESTAMP NOT NULL,
    updated_date TIMESTAMP NOT NULL
)

-- Create sizes table
/*
    Table Description: To store the different sizes
*/

CREATE TABLE sizes (
    size_id SERIAL PRIMARY KEY,
    size_name VARCHAR(5) NOT NULL,
    size_quantity INT DEFAULT 0,
    added_date TIMESTAMP NOT NULL,
    updated_date TIMESTAMP NOT NULL
)

-- Create product_sizes table
/*
    Table Description: To create a Many-To-Many relationship between 'products' and 'sizes' table.
    Composite key: product_id and size_id
*/

CREATE TABLE product_sizes (
    product_id INT REFERENCES products(product_id) ON DELETE CASCADE
    size_id INT REFERENCES sizes(size_id) ON DELETE CASCADE
    PRIMARY KEY (product_id, size_id)  -- composite key
)

-- Create product_user table
/*
    Table Description: To find out which users are part of which organisations and find out who added a specific product to the inventory.
*/

CREATE TABLE product_user (
    product_id INT REFERENCES products(product_id) ON DELETE CASCADE,
    user_id INT REFERENCES accounts(user_id) ON DELETE CASCADE,
    company_name VARCHAR(255) NOT NULL,
    added_date TIMESTAMP NOT NULL,
    updated_date TIMESTAMP NOT NULL,
    UNIQUE (product_id, user_id, company_name)
)


