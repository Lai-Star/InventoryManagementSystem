CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    product_description TEXT,
    product_sku VARCHAR(50) NOT NULL,
    product_cost DECIMAL(10,2),
    added_date TIMESTAMP NOT NULL,
    updated_date TIMESTAMP NOT NULL
)

CREATE TABLE sizes (
    size_id SERIAL PRIMARY KEY,
    size_name VARCHAR(5) NOT NULL,
    added_date TIMESTAMP NOT NULL,
    updated_date TIMESTAMP NOT NULL
)

CREATE TABLE product_sizes_mapping (
    product_id INT REFERENCES products(product_id) ON DELETE CASCADE,
    size_id INT REFERENCES sizes(size_id) ON DELETE CASCADE,
    size_quantity INT DEFAULT 0,
    PRIMARY KEY (product_id, size_id)  -- composite key
)

CREATE TABLE product_user_organisation_mapping (
    product_id INT REFERENCES products(product_id) ON DELETE CASCADE,
    user_id INT REFERENCES accounts(user_id) ON DELETE CASCADE,
    colour_id INT REFERENCES colours(colour_id) ON DELETE CASCADE,
    category_id INT REFERENCES categories(category_id) ON DELETE CASCADE,
    brand_id INT REFERENCES brands(brand_id) ON DELETE CASCADE,
    organisation_id INT REFERENCES organisations(organisation_id) ON DELETE CASCADE,
    added_date TIMESTAMP NOT NULL,
    updated_date TIMESTAMP NOT NULL,
    PRIMARY KEY (product_id, user_id)
)

CREATE TABLE colours (
    colour_id SERIAL PRIMARY KEY,
    colour_name VARCHAR(60) NOT NULL,
    added_date TIMESTAMP NOT NULL,
    updated_date TIMESTAMP NOT NULL
)

CREATE TABLE categories (
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(60) NOT NULL,
    added_date TIMESTAMP NOT NULL,
    updated_date TIMESTAMP NOT NULL
)

CREATE TABLE brands (
    brand_id SERIAL PRIMARY KEY,
    brand_name VARCHAR(60) NOT NULL,
    added_date TIMESTAMP NOT NULL,
    updated_date TIMESTAMP NOT NULL
)