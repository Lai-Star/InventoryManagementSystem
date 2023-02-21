-- Table `users`: Stores the user accounts details.
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    is_active SMALLINT DEFAULT 1,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

-- Table `organisations`: To store the different organisations that register with IMS.
CREATE TABLE organisations (
	organisation_id SERIAL PRIMARY KEY,
	organisation_name VARCHAR(255) NOT NULL,
	added_date TIMESTAMP DEFAULT NOW(),
	updated_date TIMESTAMP DEFAULT NOW()
);

-- Table `user_groups`: To store the user groups for the differnet users.
CREATE TABLE user_groups (
	user_group_id SERIAL PRIMARY KEY,
    user_group VARCHAR(255) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE user_organisation_mapping (
    user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    organisation_id INT NOT NULL REFERENCES organisations(organisation_id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, organisation_id)
);

CREATE TABLE user_group_mapping (
    user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    user_group_id INT NOT NULL REFERENCES user_groups(user_group_id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, user_group_id)
);

-- Insert default user group into user_groups table
INSERT INTO user_groups (user_group, description) VALUES ('InvenNexus User', 'Regular user using the InvenNexus application who can access all the functionalities of operations and financial analyst.');
INSERT INTO organisations (organisation_name) VALUES ('InvenNexus');

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO leon;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO leon;

SELECT * FROM users;
SELECT * FROM user_groups;
SELECT * FROM organisations;
SELECT * FROM user_group_mapping;
SELECT * FROM user_organisation_mapping;