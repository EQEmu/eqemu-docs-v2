# Merchant Data Buckets
* Merchant data buckets allow Server Operators to limit the items a merchant will show based on different data bucket comparisons and the bucket name and value provided.

## Using Merchant Data Buckets
* To use a Merchant Data Bucket you will need to set the `bucket_name`, `bucket_value`, and `bucket_comparison` fields in the `merchantlist` table for your intended item.

## How the Data Buckets are Used
* The Merchant grabs all of your Data Buckets when you open the Merchant window and then compares your values to the field values noted above.

## Comparison Types
* Merchant Data Buckets have 10 different comparison types.

## Standard Comparison Types
* There are six standard comparison types.
    - `0 (==)` checks your value and the Merchant value for `equality`
    - `1 (!=)` checks your value and the Merchant value for `inequality`
    - `2 (>=)` checks if your value is `greater than or equal to` the Merchant value
    - `3 (<=)` checks if your value is `lesser than or equal to` the Merchant value
    - `4 (>)` checks if your value is `greater than` the Merchant value
    - `5 (<)` checks if your value is `lesser than` the Merchant value

## Special Comparsion Types
* There are four special comparison types.
* These values are pipe `|` separated.
    - `6 (Any)` checks if your value is `any` of the pipe-separated Merchant values `(Example: 1|2|3)`
    - `7 (Not Any)` checks if your value is `not any` of the pipe-separated Merchant values `(Example: 4|5|6)`
    - `8 (Between)` checks if your value is `between` two pipe-separated Merchant values `(Example: 1|3)`
    - `9 (Not Between)` checks if your value is `not between` two pipe-separated Merchant values `(Example: 4|6)`
