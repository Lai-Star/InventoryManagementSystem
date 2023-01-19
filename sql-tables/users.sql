-- Create accounts table
/*
    Table Description: Stores the user accounts details.
*/

CREATE TABLE accounts (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    is_active SMALLINT DEFAULT 1,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

-- Create user_group_mapping table
/*
    Table Description: To identify which user belongs to which user group
    Composite key: user_id and user_group_id
*/

CREATE TABLE user_group_mapping (
    user_id INT NOT NULL REFERENCES accounts(user_id) ON DELETE CASCADE,
    user_group_id INT NOT NULL REFERENCES allowed_groups (user_group_id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, user_group_id)
);


-- Create allowed_groups table
/*
    Table Description: To create user groups
    Unique Constraint: user_group (to ensure that there are no duplicates inserted into that column).
*/

CREATE TABLE allowed_groups (
    user_group_id SERIAL PRIMARY KEY,
    user_group VARCHAR(255) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    added_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW()
);

