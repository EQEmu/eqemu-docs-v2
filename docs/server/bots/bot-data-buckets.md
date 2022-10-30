# Bot Data Buckets
* Bot data buckets allow Server Operators to limit the spells a Bot will cast based on different data bucket comparisons and the bucket name and value provided.

## Using Bot Data Buckets
* To use a Bot Data Bucket you will need to set the `bucket_name`, `bucket_value`, and `bucket_comparison` fields in the `bot_spells_entries` table for your intended spell.

## How the Data Buckets are Used
* Data Buckets used by the Bot use a `bucket_key-bucket_name` format.
* Meaning if your character ID is `7` and the bucket name is `test` your Data Bucket Key will be `character-7`.
* If the Data Bucket's name is `test` the Bot will attempt to find a Data Bucket called `character-7-test`.
* The Bot grabs all of your Data Buckets when spawned and then compares your values to the field values noted above.

## Comparison Types
* Bot Data Buckets have 10 different comparison types.

## Standard Comparison Types
* There are six standard comparison types.
    - `0 (==)` checks your value and the Bot value for `equality`
    - `1 (!=)` checks your value and the Bot value for `inequality`
    - `2 (>=)` checks if your value is `greater than or equal to` the Bot value
    - `3 (<=)` checks if your value is `lesser than or equal to` the Bot value
    - `4 (>)` checks if your value is `greater than` the Bot value
    - `5 (<)` checks if your value is `lesser than` the Bot value

## Special Comparsion Types
* There are four special comparison types.
* These values are pipe `|` separated.
    - `6 (Any)` checks if your value is `any` of the pipe-separated Bot values `(Example: 1|2|3)`
    - `7 (Not Any)` checks if your value is `not any` of the pipe-separated Bot values `(Example: 4|5|6)`
    - `8 (Between)` checks if your value is `between` two pipe-separated Bot values `(Example: 1|3)`
    - `9 (Not Between)` checks if your value is `not between` two pipe-separated Bot values `(Example: 4|6)`
