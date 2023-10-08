# MySQL Prepared Statements

This document contains implementation notes and information for developers that may intend to use the prepared statement C++ API for MySQL queries.

## Thread Safety

Ideally threads would have their own database connections but the current zone connections may be used concurrently in other threads. The `mysql::PreparedStmt` class locks the `DBcore` connection mutex in calls to `Prepare`, `Execute`, and when `~PreparedStmt` destruction closes the stmt in an attempt to make it safe to use with server connections.

The [MySQL C API documentation](https://dev.mysql.com/doc/c-api/8.0/en/c-api-threaded-clients.html) provides no specific guidance for using prepared statements on a connection shared by multiple threads. The only real guarantee for safety is to synchronize the connection for the lifetime of the `mysql::PreparedStmt`.

!!! warning

    Prepared statements should only be used on the main thread even if the connection is synchronized. This is because server logging is performed internally which is not synchronized.

!!! info

    The #hotfix command may be the only real use of querying in another thread.

Regular queries are synchronized by locking the connection mutex for the query until the results are buffered (`mysql_real_query` and `mysql_store_result`). `mysql::PreparedStmt` is designed to prepare, execute (possibly many times), fetch, and close separately. The following are only observed guidelines for making these safe to use on connections shared with other threads and may not be accurate:

- Synchronizing with the `DBcore` connection mutex must be done before any use of `Prepare()` or `Execute()`. The lock can be released between these calls.

- The connection mutex needs to be locked when `PreparedStmt` closes its statement during destruction (`mysql_stmt_close`).

- If the `buffer_results` option is enabled then `Fetch()` should be safe to use without locking the connection since it accesses results stored on the client.

- If result buffering is disabled then the connection cannot be used after `Execute` until all rows are fetched or freed from the result set. Failure to do so will cause a `Commands out of sync` error if another query is attempted on the same connection. **This makes it impossible to use prepared statements safely without buffering results when used with multithreaded server connections.**

    !!! note
        Unbuffered prepared statements would be possible to use safely if the connection mutex was exposed in the `DBcore` public API but the caller would be responsible for locking it before execute and holding it until all rows were fetched. The connection mutex is recursive so there would be no risk of deadlock from internal `PreparedStmt` functions also locking it. Alternatively `PreparedStmt` could be given an option on construction to hold a lock on the connection mutex for its lifetime.

## Error Handling

Prepared statements will throw a `std::runtime_error` exception for errors. This differs from regular server queries which return an empty object and requires checking for success. This means usage must be wrapped in try/catch but simplifies error handling internally.

## Options

The `StmtOptions` class contains options that may be changed to increase performance but care must be taken if the database connection is used concurrently.

`buffer_results`

:   Enabled by default. Disabling may improve performance.

    Stores the entire result set on client by calling [`mysql_stmt_store_result`](https://dev.mysql.com/doc/c-api/8.0/en/mysql-stmt-store-result.html) after executing a prepared statement.

    This transfers the full result set over the network and uses more memory but gives access to total row count and `max_length` of columns.

    If disabled memory usage is reduced and each row is fetched over the network. Disabling may also require output buffers for strings to be re-allocated while fetching since `max_length` will be unavailable.

    !!! warning
        This option must be enabled if other queries could occur on the same connection before all results are fetched or freed. Failure to do so will cause a `Commands out of sync` error.

`use_max_length`

:   Enabled by default. Requires `buffer_results` to be enabled. Disabling may improve performance when storing results.

    This causes the `max_length` of fields to be calculated when a prepared statement is executed so output buffers for strings may be pre-allocated.

    If disabled, output buffers for string columns may need to be re-allocated while fetching.

    See note in official [`mysql_stmt_store_result`](https://dev.mysql.com/doc/c-api/8.0/en/mysql-stmt-store-result.html) documentation for more information.

## Column Values

`mysql::StmtColumn` getters return a `std::optional<T>` to make it easier to handle columns that might return a `NULL` value. If the caller knows the column cannot be null, then it may just dereference the optional to get the value without the overhead of checking. Otherwise it should check if the optional has a value or use `value_or()` to return a fallback default value.

## Example

```cpp
#include <../common/mysql_stmt.h>

void foo()
{
  try
  {
    mysql::PreparedStmt stmt = content_db.Prepare("select * from spells_new where id = ? or Name = ?");
    mysql::StmtResult result = stmt.Execute({ 100, "Illusion: Feir'Dal" });
    int total_rows = result.RowCount();
    while (mysql::StmtRow row = stmt.Fetch())
    {
      int32_t id = *row.Get<int32_t>(0); // get id by col index, we know this can't be NULL so dereference
      std::string name = row["name"].value_or(""); // get str by field name, value may be NULL
    }
  }
  catch (const std::exception& ex)
  {
    // handle failure, the error is already logged with LogMySQLError
  }
}
```

## API

### DBcore::Prepare
---

`Prepare(string query)`
:   Prepares a statement and returns a [`mysql::PreparedStmt`](#mysqlpreparedstmt) object.

    Throws `std::runtime_error` if an error occurs.

### mysql::PreparedStmt
---

Types

`param_t`
:   `std::variant` of supported argument types for MySQL prepared statement execution
    ```cpp
    using param_t = std::variant<int8_t, uint8_t, int16_t, uint16_t, int32_t, uint32_t, int64_t, uint64_t, float, double, bool, std::string_view, std::nullptr_t>;
    ```

Functions

`Execute()`
`Execute(const std::vector<T>& inputs)`
`Execute(const std::vector<param_t>& inputs)`
:   Executes the prepared statement with specified input arguments and returns a [`mysql::StmtResult`](#mysqlstmtresult) object.

    Throws `std::runtime_error` if an error occurs.

`Fetch()`
:   Fetches the next row and returns a [`mysql::StmtRow`](#mysqlstmtrow) object. The returned object will evaluate as `false` when there are no more rows to fetch.

    Throws `std::runtime_error` if an error occurs.

### mysql::StmtResult
---
Represents the result of executing a prepared statement. Holds a non-owning view of column data and should  not be used if the prepared statement is executed again.

`ColumnCount()`
:   Returns the number of fields in a result set for a `SELECT` query.

`LastInsertID()`
:   Returns the last insert id of an `AUTO_INCREMENT` column generated by an `INSERT` query.

`RowCount()`
:   Returns the number of rows in a result set for a `SELECT` query. If buffering is disabled for the prepared statement then this will return 0.

`RowsAffected()`
:   Returns the number of rows affected by an `UPDATE`, `DELETE`, or `INSERT` query.

### mysql::StmtRow
---
Provides a non-owning view of column values in a result set. This object overrides `operator bool()` to evaluate as false if it does not contain a fetched row.

`ColumnCount()`
:   Returns number of fields in the row.

`GetColumn<T>(size_t index)`
`GetColumn<T>(std::string_view name)`
:   Returns a [`mysql::StmtColumn`](#mysqlstmtcolumn) pointer to the column. Returns `nullptr` if the column doesn't exist.

`Get<T>(size_t index)`
`Get<T>(std::string_view name)`
:   Returns a `std::optional<T>` with the column value as a numeric `T`. Returns `std::nullopt` if the column doesn't exist, the value is NULL, or the field type is unsupported. See `StmtColumn::Get<T>()`.

`GetStr(size_t index)`
`GetStr(std::string_view name)`
:   Returns a `std::optional<std::string>` with the column value as a string. Returns `std::nullopt` if the column doesn't exist, the value is NULL, or the field type is unsupported. See `StmtColumn::GetStr()`.

### mysql::StmtColumn
---
Stores a column value buffer.

`Get<T>()`
:   Returns a `std::optional<T>` with the column value as a numeric `T`. Returns `std::nullopt` if the column value is `NULL` or the field type is unsupported.

    !!! warning
        This will truncate number values to the specified `T` type so the caller is responsible for requesting the proper type based on database field type.

    Time and Date fields are converted to `time_t` before casting to the desired `T` type.

    Bit fields are converted to `uint64_t` before casting to the desired `T` type.

    If `T` is `bool` then string field types return `false` for empty strings and `true` otherwise.

    If `T` is not `bool` then string field types return `0` (`T` is zero-initialized) if the string cannot be converted to a number.

`GetBuf()`
:   Returns a `std::span<const uint8_t>` view of the column value buffer. This may be used to get the raw bytes and length of the value without any conversions.

`GetStr()`
:   Returns a `std::optional<std::string>` with the column value as a string. Returns `std::nullopt` if the column value is `NULL` or the field type is unsupported.

    Time and Date fields are formatted as strings specified in the [MySQL C API](https://dev.mysql.com/doc/refman/8.0/en/datetime.html).

    Bit fields are formatted as a `uint64_t` converted to a string.

`GetStrView()`
:   Returns a `std::optional<std::string_view>` view of the column string value. Returns `std::nullopt` if the column value is `NULL` or not a string. This may be used to avoid making a copy of strings if callers do not need to store it.

`Index()`
:   Returns the index of the column in the result set.

`IsNull()`
:   Returns bool specifying if the column value is `NULL`.

`IsUnsigned()`
:   Returns bool specifying if the column field type is unsigned.

`Name()`
:   Returns the field name.

`Type()`
:   Returns the [`buffer_type`](https://dev.mysql.com/doc/c-api/8.0/en/c-api-prepared-statement-type-codes.html) of the field.
