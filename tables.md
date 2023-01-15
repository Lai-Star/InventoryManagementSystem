# PostgreSQL Tables

### Accounts Table

| Column Name  | Data Type    | Constraints           |
| ------------ | ------------ | --------------------- |
| user_id      | SERIAL       | NOT NULL, PRIMARY KEY |
| username     | VARCHAR(255) | NOT NULL              |
| password     | VARCHAR(255) | NOT NULL              |
| email        | VARCHAR(255) | NOT NULL              |
| user_group   | VARCHAR(255) | NOT NULL              |
| company_name | VARCHAR(255) | NOT NULL              |
| is_active    | SMALLINT     | NOT NULL              |
| added_date   | DATE         | NOT NULL              |
| updated_date | DATE         | NOT NULL              |

### Products Table

| Column Name         | Data Type    | Constraints |
| ------------------- | ------------ | ----------- |
| product_id          | SERIAL       | PRIMARY KEY |
| product_name        | VARCHAR(255) | NOT NULL    |
| product_description | TEXT         | NOT NULL    |
| product_sku         | VARCHAR(255) | NOT NULL    |
| product_colour      | VARCHAR(255) | NOT NULL    |
| XXS                 | INTEGER      |             |
| XS                  | INTEGER      |             |
| S                   | INTEGER      |             |
| M                   | INTEGER      |             |
| L                   | INTEGER      |             |
| XL                  | INTEGER      |             |
| XXL                 | INTEGER      |             |
| product_quantity    | INTEGER      | NOT NULL    |

### Orders Table

| Column Name    | Data Type    | Constraints |
| -------------- | ------------ | ----------- |
| order_id       | INT          | PRIMARY KEY |
| product_id     | INT          | FOREIGN KEY |
| order_quantity | INT          | NOT NULL    |
| order_date     | DATE         | NOT NULL    |
| order_status   | VARCHAR(255) | NOT NULL    |

### Customers Table

| Column Name      | Data Type    | Constraints |
| ---------------- | ------------ | ----------- |
| customer_id      | INT          | PRIMARY KEY |
| customer_name    | VARCHAR(255) | NOT NULL    |
| customer_address | VARCHAR(255) | NOT NULL    |
| customer_phone   | VARCHAR(255) | NOT NULL    |
| customer_email   | VARCHAR(255) | NOT NULL    |

### Sales Table

| Column Name   | Data Type     | Constraints |
| ------------- | ------------- | ----------- |
| sale_id       | INT           | PRIMARY KEY |
| product_id    | INT           | FOREIGN KEY |
| customer_id   | INT           | FOREIGN KEY |
| sale_quantity | INT           | NOT NULL    |
| sale_date     | DATE          | NOT NULL    |
| sale_price    | DECIMAL(10,2) | NOT NULL    |

### Inventory Table

| Column Name        | Data Type | Constraints |
| ------------------ | --------- | ----------- |
| inventory_id       | INT       | PRIMARY KEY |
| product_id         | INT       | FOREIGN KEY |
| inventory_quantity | INT       | NOT NULL    |
| inventory_date     | DATE      | NOT NULL    |

### Transactions Table

| Column Name          | Data Type    | Constraints |                              |
| -------------------- | ------------ | ----------- | ---------------------------- |
| transaction_id       | INT          | PRIMARY KEY |                              |
| product_id           | INT          | FOREIGN KEY |                              |
| transaction_quantity | INT          | NOT NULL    |                              |
| transaction_type     | VARCHAR(255) | NOT NULL    | (e.g purchase, sale, return) |
| transaction_date     | DATE         | NOT NULL    |                              |
