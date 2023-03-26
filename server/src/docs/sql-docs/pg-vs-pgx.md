# `pgx` vs `pq` PostgreSQL driver

- [pq driver](https://github.com/lib/pq)
- [pgx driver](https://github.com/jackc/pgx)

## `pq`

- `pq` implements the standard `sql.DB` specification, and as such, may be the first choice that most people new to the language go for.
- If you are looking for a driver, which simply gives you access to a standard relational database, `pq` is just fine.
- Unfortunately, `pq` has not been in active feature development as of a couple of years.
- It is still under active maintenance, so anything critical will still make it into the repo, but it may eventually diverge from the development of future PostgreSQL versions.

# `pgx`

1. Performance

- `pgx` uses a binary protocol for communication with PostgreSQL which reduces the overhead of marshalling and unmarshalling data.
- `pgx` provides a higher level of control over the database connection pool and supports connection pooling mechanisms that can reduce the latency and overhead associated with establishing a connection to the database.

```go
import (
	_ "github.com/lib/pq" // Don't forget the driver!
)

dbPath := "postgres://path_to_the_db"
db, err := sql.Open("postgres", dbPath)
```

2. Extensibility

- `pgx` provides a low-level interface that allows for greater flexibility in implementing custom solutions and integrating with third-party libraries.
- `pgx` supports custom types and composite types in a more flexible manner than the pq package.

3. Type Safety

- `pgx` provides more type safety by using a parameterized query API that reduces the risk of SQL injection attacks and improves the readability and maintainability of the code.

4. Richer Feature Set

- `pgx` provides support for asynchronous queries, prepared statements, and named transactions.
- `pgx` supports multiple result sets, bulk data transfer, and binary copy modes.

5. Better maintainability

- `pgx` has a cleaner and more modular codebase than the pq package, which makes it easier to maintain and extend over time.
- `pgx` provides more consistent error handling and logging mechanisms that can help developers quickly identify and troubleshoot issues.

```go
import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name string
	var weight int64
	err = m.DB.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, weight)
}
```
