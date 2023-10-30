# Database Schema Migrations

## Adding New Required Schema Migrations

It's extremely simple

## Changes in the Source

You need to increment a define in the source that dictates what database version the source SHOULD be running at

**Location** `common/version.h`

The database version will need to match the manifest entry you have added, more on that in a moment

`CURRENT_BINARY_DATABASE_VERSION = 9240`

## The Manifest

The manifest is located at **common/database/database_update_manifest.cpp**

Add a struct representing your migration 

```cpp
	ManifestEntry{
		.version = 9240,
		.description = "2023_10_29_variables_id.sql",
		.check = "SHOW COLUMNS FROM `variables` LIKE 'id'",
		.condition = "empty",
		.match = "",
		.sql = R"(
ALTER TABLE `variables`
ADD COLUMN `id` int(11) NOT NULL AUTO_INCREMENT FIRST,
DROP PRIMARY KEY,
ADD PRIMARY KEY (`id`) USING BTREE,
ADD UNIQUE INDEX(`varname`);
)"
	},
```

That's it! As far as what is needed from a developer to have the server run the migration, that is all you need to do.

You can test it by running world manually after you compile. Please test your database migrations before submitting a PR, it's a very simple mistake to avoid trying to fix later.

## Manifest Conditions

```cpp
// see struct definitions for what each field does
// struct ManifestEntry {
// 	int         version{};     // database version of the migration
// 	std::string description{}; // description of the migration ex: "add_new_table" or "add_index_to_table"
// 	std::string check{};       // query that checks against the condition
// 	std::string condition{};   // condition or "match_type" - Possible values [contains|match|missing|empty|not_empty]
// 	std::string match{};       // match field that is not always used, but works in conjunction with "condition" values [missing|match|contains]
// 	std::string sql{};         // the SQL DDL that gets ran when the condition is true
// };
```
