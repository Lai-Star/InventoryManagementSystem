CREATE TABLE IF NOT EXISTS products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    product_description TEXT,
    product_sku VARCHAR(50) NOT NULL,
    product_cost DECIMAL(10,2),
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_sizes (
    user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    size_id SERIAL PRIMARY KEY,
    size_name VARCHAR(5) NOT NULL,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS organisation_sizes (
    organisation_id INT REFERENCES organisations(organisation_id) ON DELETE CASCADE,
    size_id SERIAL PRIMARY KEY,
    size_name VARCHAR(5) NOT NULL,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_product_sizes_mapping (
    product_id INT REFERENCES products(product_id) ON DELETE CASCADE,
    size_id INT REFERENCES user_sizes(size_id) ON DELETE CASCADE,
    size_quantity INT DEFAULT 0,
    PRIMARY KEY (product_id, size_id)  -- composite key
);

CREATE TABLE IF NOT EXISTS organisation_product_sizes_mapping (
    product_id INT REFERENCES products(product_id) ON DELETE CASCADE,
    size_id INT REFERENCES organisation_sizes(size_id) ON DELETE CASCADE,
    size_quantity INT DEFAULT 0,
    PRIMARY KEY (product_id, size_id)  -- composite key
);

CREATE TABLE IF NOT EXISTS user_categories (
    user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(60) NOT NULL,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS organisation_categories (
    organisation_id INT REFERENCES organisations(organisation_id) ON DELETE CASCADE,
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(60) NOT NULL,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_colours (
    user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    colour_id SERIAL PRIMARY KEY,
    colour_name VARCHAR(60) NOT NULL,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS organisation_colours (
    organisation_id INT REFERENCES organisations(organisation_id) ON DELETE CASCADE,
    colour_id SERIAL PRIMARY KEY,
    colour_name VARCHAR(60) NOT NULL,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_brands (
    user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    brand_id SERIAL PRIMARY KEY,
    brand_name VARCHAR(60) NOT NULL,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS organisation_brands (
    organisation_id INT REFERENCES organisations(organisation_id) ON DELETE CASCADE,
    brand_id SERIAL PRIMARY KEY,
    brand_name VARCHAR(60) NOT NULL,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS product_user_mapping (
    product_id INT REFERENCES products(product_id) ON DELETE CASCADE,
    user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    colour_id INT REFERENCES user_colours(colour_id) ON DELETE CASCADE,
    category_id INT REFERENCES user_categories(category_id) ON DELETE CASCADE,
    brand_id INT REFERENCES user_brands(brand_id) ON DELETE CASCADE,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY (product_id)
);

CREATE TABLE IF NOT EXISTS product_organisation_mapping (
    product_id INT REFERENCES products(product_id) ON DELETE CASCADE,
    colour_id INT REFERENCES organisation_colours(colour_id) ON DELETE CASCADE,
    category_id INT REFERENCES organisation_categories(category_id) ON DELETE CASCADE,
    brand_id INT REFERENCES organisation_brands(brand_id) ON DELETE CASCADE,
    organisation_id INT REFERENCES organisations(organisation_id) ON DELETE CASCADE,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY (product_id)
);

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO leon;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO leon;

TRUNCATE TABLE products CASCADE;
TRUNCATE TABLE user_product_sizes_mapping CASCADE;
TRUNCATE TABLE organisation_product_sizes_mapping CASCADE;
TRUNCATE TABLE product_user_mapping;
TRUNCATE TABLE product_organisation_mapping;

TRUNCATE TABLE user_sizes CASCADE;
TRUNCATE TABLE organisation_sizes CASCADE;
TRUNCATE TABLE user_categories CASCADE;
TRUNCATE TABLE organisation_categories CASCADE;
TRUNCATE TABLE user_colours CASCADE;
TRUNCATE TABLE organisation_colours CASCADE;
TRUNCATE TABLE user_brands CASCADE;
TRUNCATE TABLE organisation_brands CASCADE;