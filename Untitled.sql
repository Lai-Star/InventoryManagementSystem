-- Create user table

CREATE TABLE IF NOT EXISTS users (
	user_id SERIAL PRIMARY KEY,
	username VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	email VARCHAR(255),
	isActive INT NOT NULL,
	added_date DATE NOT NULL,
	updated_date DATE NOT NULL
);

SELECT * FROM users;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO leon;

INSERT INTO users (username, password, email, isActive, added_date, updated_date) VALUES ("lowjiewei","password","test",1,now(),now());