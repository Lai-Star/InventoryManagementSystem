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

|    User Group     |                                                    Description                                                     |
| :---------------: | :----------------------------------------------------------------------------------------------------------------: |
|       Admin       |         Managing user accounts and permissions. Implementing new features and improvements to the system.          |
|     IMS User      |              Can access all the functionalities of operations and financial analyst. (Self-Employed)               |
|    Operations     |  Use IMS to manage and track sales, customer details, inventory and place orders for the business. (Organisation)  |
| Financial Analyst | Use IMS to analyze financial data, such as profit margins, monthly profits, and inventory turnover. (Organisation) |

## User Stories

- As a user, I want to be able to login into my IMS account with a secure login method (2FA Authentication).
- As a non-user, I want to be able to sign up for a IMS account in the login page.
- As an admin, I want to able to create accounts for users and manage the permissions.
- As an admin, I want to able to view, edit and delete the accounts in the IMS.
- As a user, I want to be able to add a new product to the IMS.
- As a user, I want to be able to view, edit and delete the products in my IMS account.

## Entity-Relationship Diagram (ERD)

#### User Management ERD

<img src='./server/src/docs/ERD-User-Management.png' alt='User Management ERD Diagram'>

#### Inventory Management ERD

<img src='./server/src/docs/ERD-Inventory-Management.png' alt='Inventory Management ERD Diagram'>

## 2FA Authentication

- 2FA provides an additional layer of security as it mitigates the risks associated with weak or stolen passwords.
- By requiring the user to provide 2 different authentication factors, 2FA helps to protect against unauthorized access to an account or system.
- Cryptographically secure pseudorandom number generators (CSPRNG) were used to generate one-time passwords (OTPs) for two-factor authentication (2FA) over Psedorandom Number Generators (PRNG) due to the **increased security, enhanced randomness, and improved reliability**.
- In the **Generate2FA()** function, the OTP was generated in the format of 5 alphabets (lower and upper case) and 6 numbers separated by a hyphen "-" to provide a high level of randomness.

## Unit Testing - Test Driven Development (TDD)

- The main purpose of unit testing is to test the functionality of individual units of code to ensure that they are working as intended.
- Unit testing can also be used to establish a foundation for automated testing and continuous integration, which helps to maintain the quality and reliability of the software over time.
- The following steps were implemented in the process of TDD:
  1. Plan your test cases: Write down the different scenarios you want to test, and write down the expected input and output for each test case.
  2. Write the function: Write the function you want to test in code. At this point, the function should have any actual implementation. It should just have the correct function signature and return type.
  3. Write the unit test function: Write a unit test function that tests the function in step 2. The unit test function should take in the expected input and output for each test case, and use assertions to check that the function produces the correct output when given the expected input.
  4. Run the tests and refactor (if necessary).

---

- `coverage`
  - `alias coverage='go test -coverprofile=coverage.out && go tool cover -html=coverage.out'`
- Running ALL tests on /src directory
  - `go test -v ./...`
  - `go test ./...`

