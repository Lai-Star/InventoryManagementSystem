## Using `INSERT` and `SELECT` together

```sql
SQL_INSERT_INTO_USER_ORGANISATION_MAPPING = "INSERT INTO user_organisation_mapping (user_id, organisation_id) " +
                        "SELECT $1, organisation_id " +
                        "FROM organisations " +
                        "WHERE organisation_name = $2;"
```
