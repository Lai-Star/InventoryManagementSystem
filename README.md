# Inventory Management System (IMS)

- An inventory management system (IMS) is crucial for businesses that need to efficiently track and manage their inventory. (primarily working on the IMS for my own online business).

## Objectives

- Maintain quality control through accurate tracking of transactions involving consumer goods.
- Use sales data to constantly update and maintain precise inventory records.
- Avoid stock shortages by alerting the wholesaler when it is time to restock (through Email SMTP).
- Reduce errors in stock recording through automation.

## 3-Tier Architecture

|     Layer      |  Software  |
| :------------: | :--------: |
|    Database    | PostgreSQL |
|     Logic      |   Golang   |
| Presentational |  ReactJS   |

## User Group

|       User Group       |                                            Description                                            |
| :--------------------: | :-----------------------------------------------------------------------------------------------: |
|         Admin          | Managing user accounts and permissions. Implementing new features and improvements to the system. |
| Retail business owners |                                Use IMS to manage their inventory.                                 |
| Retail store employees |                     Use IMS to check stock levels and place orders as needed.                     |
|   Warehouse managers   |          Use IMS to track inventory for distribution centers or fulfillment warehouses.           |

## User Stories

- As a user, I want to be able to login into my IMS account with a secure login method (2FA Authentication).
- As a non-user, I want to be able to sign up for a IMS account in the login page.
- As an admin, I want to able to create accounts for users and manage the permissions.
- As an admin, I want to able to view, edit and delete the accounts in the IMS.
- As a user, I want to be able to add a new product to the IMS.
- As a user, I want to be able to view, edit and delete the products in my IMS account.

## Test Driven Development (TDD)

- `coverage`
  - `alias coverage='go test -coverprofile=coverage.out && go tool cover -html=coverage.out'`
