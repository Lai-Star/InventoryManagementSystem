CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    product_description TEXT,
    product_sku VARCHAR(50) NOT NULL,
    product_cost DECIMAL(10,2),
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE organisation_sizes (
    size_id SERIAL PRIMARY KEY,
    organisation_id INT NOT NULL,
    size_name VARCHAR(5) NOT NULL,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE user_sizes (
    size_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    size_name VARCHAR(5) NOT NULL,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE product_sizes_mapping (
    product_id INT REFERENCES products(product_id) ON DELETE CASCADE,
    size_id INT REFERENCES sizes(size_id) ON DELETE CASCADE,
    size_quantity INT DEFAULT 0,
    PRIMARY KEY (product_id, size_id)  -- composite key
);

CREATE TABLE user_categories (
    user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(60) NOT NULL,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE organisation_categories (
    organisation_id INT REFERENCES organisations(organisation_id) ON DELETE CASCADE,
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(60) NOT NULL,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE user_colours (
    user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    colour_id SERIAL PRIMARY KEY,
    colour_name VARCHAR(60) NOT NULL,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE organisation_colours (
    organisation_id INT REFERENCES organisations(organisation_id) ON DELETE CASCADE,
    colour_id SERIAL PRIMARY KEY,
    colour_name VARCHAR(60) NOT NULL,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE user_brands (
    user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    brand_id SERIAL PRIMARY KEY,
    brand_name VARCHAR(60) NOT NULL,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE organisation_brands (
    organisation_id INT REFERENCES organisations(organisation_id) ON DELETE CASCADE,
    brand_id SERIAL PRIMARY KEY,
    brand_name VARCHAR(60) NOT NULL,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE product_user_mapping (
    product_id INT REFERENCES products(product_id) ON DELETE CASCADE,
    user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    colour_id INT REFERENCES colours(colour_id) ON DELETE CASCADE,
    category_id INT REFERENCES categories(category_id) ON DELETE CASCADE,
    brand_id INT REFERENCES brands(brand_id) ON DELETE CASCADE,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY (product_id)
);

CREATE TABLE product_organisation_mapping (
    product_id INT REFERENCES products(product_id) ON DELETE CASCADE,
    colour_id INT REFERENCES colours(colour_id) ON DELETE CASCADE,
    category_id INT REFERENCES categories(category_id) ON DELETE CASCADE,
    brand_id INT REFERENCES brands(brand_id) ON DELETE CASCADE,
    organisation_id INT REFERENCES organisations(organisation_id) ON DELETE CASCADE,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY (product_id)
);

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO leon;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO leon;