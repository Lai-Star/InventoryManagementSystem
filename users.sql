-- Create user table

CREATE TABLE accounts (
	user_id SERIAL NOT NULL PRIMARY KEY,
	username VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	user_group VARCHAR(255) NOT NULL,
	company_name VARCHAR(255) NOT NULL,
	is_active SMALLINT NOT NULL,
	added_date DATE NOT NULL,
	updated_date DATE NOT NULL
);

SELECT * FROM accounts;

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO leon;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO leon;

-- Create products table

CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    product_description TEXT NOT NULL,
    product_sku VARCHAR(255) NOT NULL,
    product_colour VARCHAR(255) NOT NULL,
	XXS INTEGER,
    XS INTEGER,
    S INTEGER,
    M INTEGER,
    L INTEGER,
    XL INTEGER,
    XXL INTEGER,
    product_quantity INTEGER NOT NULL
);

