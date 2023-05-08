# Issues faced

1. Dependency Injection

2. Not using pointers

3. pgx.Connect does not allow `Query`. Have to use pgxpool.Connect

4. Adding `defer rows.Close()`
   - Issue happens when querying for multiple rows.
     - Previously, I did not close the rows (careless mistake). I encountered the issue of `context deadline exceeded` with my Golang context.
   - By deferring rows.Close(). This releases any resources held by the rows no matter how the function returns.
   - Looping all the way through the rows also closes it implicitly, but it is better to use defer to make sure rows is closed no matter what.
   - [Querying in Go](https://go.dev/doc/database/querying)
