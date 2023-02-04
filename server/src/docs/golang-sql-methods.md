# Golang SQL DB Method

## `db.QueryRow` vs `db.Exec`

- `db.QueryRow`
  - Used to retrieve a single row from a database query.
  - If the query returns multiple rows, only the first row is returned.
  - If the query returns a single row, that row is returned as a 'Scan-able' object (such as `sql.Row`).
  - The `Scan` method can then be used to extract the values from the returned row into variables.
- `db.Exec`
  - Used to execute a query that does not return any rows such as INSERT, UPDATE or DELETE statement.
  - The `Exec` method returns an `sql.Result` object that can be used to retrieve information about the number of affected rows, but it cannot be used to retrieve the values of the inserted row.

## `db.QueryRow` vs `db.Query`

- `db.QueryRow`
    - Returns a single row, which can be returned as a scan of the returned result. 
    - Use this when you expect to receive **exactly one row** back from the database.
- `db.Query`
    - Returns a set of rows, represented as a `*Rows` value, which can be used to iterate through the returned result set using a loop.
    - Use this when you expect to receive **0 or more rows**.