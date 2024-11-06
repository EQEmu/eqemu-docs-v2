# Database API

The database quest API allows scripts to connect to a MySQL/MariaDB database and perform queries without the need to install [LuaSQL](https://lunarmodules.github.io/luasql/) or use [perl DBI](https://dbi.perl.org/).

## Database Resources

Ownership of handles that wrap database resources are adopted by the scripting language. The internal resource will be freed when the handle is destroyed or its `close()` method is called.

Lua is garbage collected so these handles should be closed explicitly when no longer needed. If not closed then Lua will hold database resources open for an indeterminate time until its GC runs. Since errors will halt script execution this can prevent handles from being closed. See the [errors](#errors) section for a workaround.

Perl is reference counted so handles will be closed when its last reference is destroyed by going out of scope (even on errors).

## Database Connections

Connections can be made manually or by using exported constants to connect to the server databases specified in `eqemu_config`.

The API uses the current zone connections by default when using a server database constant. An overload is available to create [`Database`](#database) objects that make new connections.

Making new database connections avoids possible concurrency issues since the zone database connections might be used in other threads. The connection is synchronized for some prepared statement operations but safety isn't guaranteed. See the developer [thread safety](../developer/mysql-stmt.md#thread-safety) section for more info.

New connections also allow the use of [unbuffered](../developer/mysql-stmt.md#options) results with prepared statements. If a prepared statement is executed with buffering disabled, then all results must be fetched or freed before any other queries on that connection can occur. Failure to consume all results before the next query will cause a "Commands out of sync" error. If zone connections were used this would cause all zone queries on that connection to start failing if a script error occured before results could be fetched, or the server tried to perform a query in another thread before the script finished fetching results.

=== "Lua"

    ```lua
    local db = Database() -- connects to default database using eqemu config credentials
    local db = Database(Database.Default)
    local db = Database(Database.Content)
    local db = Database(Database.Default, false) -- uses zone's database connection
    local db = Database("127.0.0.1", "user", "pass", "db", 3306)
    ```

=== "Perl"

    ``` perl
    my $db = Database::new(); # connects to default database using eqemu config credentials
    my $db = Database::new(Database::Default);
    my $db = Database::new(Database::Content);
    my $db = Database::new(Database::Content, 0); # uses zone's content_db connection
    my $db = Database::new("127.0.0.1", "user", "pass", "db", 3306);
    ```

## Errors

Errors in prepared statements will propagate the `std::runtime_error` exception as a quest error and automatically halt the script. This is similar to the behavior of using `assert()` in Lua or `die` in perl to error out of the script.

This means db resources cannot be freed automatically on errors since the script gets halted and will never reach any calls to `close()` resources. Perl will destroy objects in scope via reference counting so it has less risk of holding resources open. In Lua this is a problem because the GC may not run immediately (if ever). To work around this in Lua, scripts can wrap API calls with `pcall()` to manually catch db errors and free resources.

#### Catching Lua Errors

  ``` lua
  local query = "select * from badtable where id = ?"
  local db = Database()

  local ok, stmt = pcall(function() return db:prepare(query) end)
  if not ok then
    db:close()
    if stmt then error("error: " .. stmt) end
  end

  local ok, err = pcall(function() return stmt:execute({1}) end)
  if not ok then
    stmt:close()
    db:close()
    if err then error("error: " .. err) end
  end

  -- do stuff

  stmt:close()
  db:close()
  ```

Manually catching errors like this should generally not be necessary except for custom handling.

## Example

!!! warning

    Query strings formed by concatenating the values can be vulnerable to SQL injection. It's recommended to always use value placeholders with prepared statements.

=== "Lua"

    ```lua
    local db = Database(Database.Content)
    local stmt = db:prepare("select * from items where id = ? or name = ?")
    stmt:execute({ 1001, "Soul's Eye" })
    local row = stmt:fetch_hash()
    while row do
      eq.debug(string.format("id %d, name %s", row.id, row.Name))
      row = stmt:fetch_hash() -- next
    end
    stmt:close()
    db:close()
    ```

=== "Perl"

    ```perl
    my $db = Database::new(Database::Content);
    my $stmt = $db->prepare("select * from items where id = ? or name = ?");
    $stmt->execute(1001, "Soul's Eye");
    while (my $row = $stmt->fetch_hashref())
    {
      quest::debug("id: " . $row->{"id"} . " name: " . $row->{"Name"});
    }
    $stmt->close();
    $db->close();
    ```

## Lua Iterator

  It may be convenient to use custom iterators to fetch rows in Lua:

  ```lua
  local function array_iter(stmt)
    return function() return stmt:fetch_array() end
  end

  local function hash_iter(stmt)
    return function() return stmt:fetch_hash() end
  end

  local stmt = db:prepare("select * from items where id = ? or name = ?")
  stmt:execute({ 1001, "Soul's Eye" })
  for row in hash_iter(stmt) do
    eq.debug(string.format("id %d, name %s", row["id"], row["Name"]))
  end
  ```

## Passing NULL

Lua tables cannot have `nil` values. Using `nil` in calls to `execute()` will not work and result in the wrong number of expected arguments. Passing `NULL` in `execute()` requires passing a table instead. Perl can use `undef`.

=== "Lua"

    ```lua
    local null = {} -- re-usable
    local stmt = db:prepare("INSERT INTO some_table (id, nullable, str) VALUES (?, ?, ?)")
    stmt:execute({ 1, null, "value" })
    eq.debug("affected: " .. stmt:rows_affected())
    ```

=== "Perl"

    ```perl
    my $stmt = $db->prepare("INSERT INTO some_table (id, nullable, str) VALUES (?, ?, ?)");
    $stmt->execute(1, undef, "value");
    quest::debug("affected: " . $stmt->rows_affected());
    ```

## BigInt Columns

Lua prior to version 5.3 only has 53-bit integer precision because all numbers are stored as `double`. 64-bit `BIGINT` column values are pushed as strings when prepared statement rows are fetched to avoid loss of precision. This leaves scripts responsible for either converting the value to a number and losing precision, using a big number library, or using [LuaJIT FFI](https://luajit.org/ext_ffi.html). Big integers larger than 2^53^ should also be passed as strings when executing prepared statements from scripts.

!!! note
    Lua 5.3 and later support a separate integer type stored as signed 64-bit integer by default. Large unsigned 64-bit integers will overflow into negative signed values in those versions.

In perl 64-bit integers are supported when building for 64-bit targets so no special handling is done when pushing from the quest API.

## API

### Constants
---

`Connection` constants are used to identify zone database connections.

=== "Lua"

    | Constant          | Value | Description
    | :---------------- | :---- | :----------------------------------------
    | Database.Default  | 0     | Reference the default database connection
    | Database.Content  |	1     |	Reference the content database connection

=== "Perl"

    | Constant          | Value | Description
    | :---------------- | :---- | :----------------------------------------
    | Database::Default | 0     | Reference the default database connection
    | Database::Content |	1     |	Reference the content database connection

### Database
---

Constructors

=== "Lua"

    `Database()`
    : Returns a `Database` object that connects to the server's default database connection. If the connection fails then a `std::runtime_error` exception is thrown and script execution halts.

    `Database(Connection type, bool connect = false)`
    : Returns a `Database` object for the specified server database type specified in the eqemu_config. If `connect` is `true` a new connection is made to the database. If `connect` is `false` the `Database` object uses the zone's current active server connection. If the connection fails or the connection type is invalid a `std::runtime_error` exception is thrown and script execution halts.

    `Database(string host, string user, string pass, string dbname, int port)`
    : Returns a `Database` object that connects to the specified database. If the connection fails then a `std::runtime_error` exception is thrown and script execution halts.

=== "Perl"

    `Database::new()`
    : Returns a `Database` object that connects to the server's default database connection. If the connection fails then a `std::runtime_error` exception is thrown and script execution halts.

    `Database::new(Connection type, bool connect = false)`
    : Returns a `Database` object for the specified server database type specified in the eqemu_config. If `connect` is `true` a new connection is made to the database. If `connect` is `false` the `Database` object uses the zone's current active server connection. If the connection fails or the connection type is invalid a `std::runtime_error` exception is thrown and script execution halts.

    `Database::new(string host, string user, string pass, string dbname, int port)`
    : Returns a `Database` object that connects to the specified database. If the connection fails then a `std::runtime_error` exception is thrown and script execution halts.

Functions

`close()`
:   Closes the database connection. The quest object is invalidated and should not be used after this is called.

`prepare(string query)`
:   Prepares a statement and returns a [`MySQLPreparedStmt`](#mysqlpreparedstmt) object. Throws `std::runtime_error` on error which halts the script with a quest error message.

### MySQLPreparedStmt
---
`close()`
:   Deletes the underlying `mysql::StmtResult` and `mysql::PreparedStmt` resources. Closes the prepared statement by calling `mysql_stmt_close()` which will free any active result set. The quest object is invalidated and should not be used after this is called.

`execute()`
`execute(array values)`
:   Executes the prepared statement with specified input arguments. Results are accessible from the current object. Invalidates any previous result set.

    !!! example
        === "Lua"
            ```lua
            local stmt = db:prepare("select * from items where id = ? or name = ?")
            stmt:execute({ 1001, "Soul's Eye" })
            eq.debug("result row count: " .. stmt:num_rows())
            ```

        === "Perl"
            ```perl
            my $stmt = $db->prepare("select * from items where id = ? or name = ?");
            $stmt->execute(1001, "Soul's Eye");
            quest::debug("result row count: " . $stmt->num_rows());
            ```

`fetch()`
: Alias for `fetch_array`.

`fetch_array()`
:   Fetches the next row and returns an array of column values. Value types will be either numbers or strings to match field types unless `NULL`.

    === "Lua"
        `NULL` column values will not exist in the returned table. This can leave gaps in table indexes which means `ipairs` should not be used to iterate it. Returns `nil` if no more rows to fetch.

    === "Perl"
        `NULL` column values will have `undef` values in the array. Returns an empty value if no more rows to fetch.


`fetch_arrayref()` <sub><sup>==Perl only==</sup></sub>
:   Fetches the next row and returns a reference to an array of column values. `NULL` column values will have `undef` values in the array. Other value types will be either numbers or strings to match field types. Returns an empty value if no more rows to fetch.

`fetch_hash()` <sub><sup>==Lua only==</sup></sub>
:   Fetches the next row and returns a hash table of column values with field name keys. `NULL` column values will not exist in the returned table. Other value types will be either numbers or strings to match field types. Returns `nil` if no more rows to fetch.

`fetch_hashref()` <sub><sup>==Perl only==</sup></sub>
:   Fetches the next row and returns a reference to a hash of column values with field name keys. `NULL` column values will have `undef` values in the hash. Other value types will be either numbers or strings to match field types. Returns an empty value if no more rows to fetch.

`insert_id()`
:   Returns the last insert id of an `AUTO_INCREMENT` column generated by an `INSERT` query.

`num_fields()`
:   Returns the number of fields in a result set for a `SELECT` query.

`num_rows()`
:   Returns the number of rows in a result set for a `SELECT` query. If buffering is disabled for the prepared statement then this will return 0.

`rows_affected()`
:   Returns the number of rows affected by an `UPDATE`, `DELETE`, or `INSERT` query.

`set_options(hash options)`
:   Sets prepared statement [options](../developer/mysql-stmt.md#options) using a hash table with option name as key. Preserves current value of unspecified options.

    | Option | Default | Description |
    | :----- | :------ | :---------- |
    | `buffer_results` | true | Enable buffering the entire result set of an executed prepared statement |
    | `use_max_length` | true | Enable the `max_length` of fields to be calculated on execution (requires buffering) |

    !!! example
        === "Lua"
            ```lua
            local stmt = db:prepare("select * from table")
            stmt:set_options({ ["buffer_results"] = false })
            ```

        === "Perl"
            ```perl
            my $stmt = $db->prepare("select * from table");
            $stmt->set_options(("buffer_results" => 0));
            ```