# SQL Prepare Statement

- Using SQL Prepare Statements help mitigate SQL injection.
- Generally recommended to use prepared statements instead of just placeholder in your database queries whenever possible, to take advantage of their benefits.

## Advantages

1. Reusability

- When you use a prepared statement, the database prepares the statement once and keeps it in memory.
- This means that if you need to execute the same query with different parameter values, you can reuse the prepared statement, which can save some overhead of parsing the query each time you execute it.

2. Performance

- Prepared statements can be faster than regular queries because the database can optimize the execution plan for the statement once and reuse it.

3. Security

- Prepared statements provide an additional layer of security by ensuring that the input parameters are properly sanitized and formatted before that are included in the query.
- This helps prevent some types of SQL injection attacks that might bypass simple placeholder-based protection.

```go
// Plain placeholder values to execute SQL query
func getUserByID(db *sql.DB, userID int) (string, error) {
    var name string
    err := db.QueryRow("SELECT name FROM users WHERE id = $1", userID).Scan(&name)
    if err != nil {
        return "", err
    }
    return name, nil
}

// Execute SQL query using Prepare Statement
func getUserByIDPrepared(db *sql.DB, userID int) (string, error) {
    var name string
    stmt, err := db.Prepare("SELECT name FROM users WHERE id = $1")
    if err != nil {
        return "", err
    }
    defer stmt.Close()
    err = stmt.QueryRow(userID).Scan(&name)
    if err != nil {
        return "", err
    }
    return name, nil
}
```

## Automatic Prepared Statement Caching

- [Automatic Prepared Statement Caching for pgx](https://github.com/jackc/pgx/wiki/Automatic-Prepared-Statement-Caching)
- Prepared statements can be manually created with the Prepare method.
- However, this is rarely necessary because pgx indicates an automatic statement cache by default.
- Queries run through normal Query, QueryRow and Exec functions are automatically prepared on first execution and the prepared statement is reused on subsequent executions.
