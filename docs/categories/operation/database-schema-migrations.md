# Database Schema Migrations

## Adding New Required Schema Migrations

* It's extremely simple
* When you add your .sql file, you still will add it as traditionally to our regular path

`utils/sql/git/required` - [Github Link](https://github.com/EQEmu/Server/tree/master/utils/sql/git/required)

## Changes in the Source

* You need to increment a define in the source that dictates what database version the source SHOULD be running at

**Location** `common/version.h`

The database version will need to match the manifest entry you have added, more on that in a moment

`CURRENT_BINARY_DATABASE_VERSION = 9058`

## The Manifest

* The manifest is stored always on Github and contains all the definitions and logic for determining if a database needs to update
  * `utils/sql/db_update_manifest.txt`
  * [Manifest Link](https://github.com/EQEmu/Server/blob/master/utils/sql/db_update_manifest.txt)

Add a line to the bottom of the file, it is going to look similar to the following

```text
9058|2014_11_17_example_update_file.sql|SHOW TABLES LIKE 'character_mercenaries'|empty|
```

* This example would then have users run `2014_11_17_example_update_file.sql` when they don't have the 'character\_mercenaries' table because of the empty condition
* That's it! As far as what is needed from a developer to have the server run the migration, that is all you need to do

## Manifest Conditions

```text
Example: Version|Filename.sql|Query_to_Check_Condition_For_Needed_Update|match type|text to match
	0 = Database Version
	1 = Filename.sql
	2 = Query_to_Check_Condition_For_Needed_Update
	3 = Match Type - If condition from match type to Value 4 is true, update will flag for needing to be ran
		contains = If query results contains text from 4th value
		match = If query results matches text from 4th value
		missing = If query result is missing text from 4th value
		empty = If the query results in no results
		not_empty = If the query is not empty
	4 = Text to match
```

## Other Manifest Examples

#### Missing

```text
9056|2014_11_08_RaidMembers.sql|SHOW COLUMNS FROM `raid_members` LIKE 'groupid'|missing|unsigned
```

This entry is looking for what the column looks like in raid\_members and to see if it is an unsigned integer

The match type is `missing`, so I'm looking to see if the word 'unsigned' is missing from the table. In this case if unsigned was missing, we need to run this update because that is what the update did

_If Missing_ If the table is missing the column, it will run the SQL file specified above `2014_11_08_RaidMembers.sql`

```text
ALTER TABLE `raid_members` CHANGE COLUMN `groupid` `groupid` INT(4) UNSIGNED NOT NULL DEFAULT '0' AFTER `charid`;
```

#### Contains

```text
9055|2014_10_30_special_abilities_null.sql|SHOW COLUMNS FROM `npc_types` LIKE 'special_abilities'|contains|NO
```

This entry is looking for what the column looks like in npc\_types, column 'special\_abilities' to see if it contains the word 'NO' \(If special\_abilities is nullable\), which if you look at the SQL result at the given time before this update is applied

You will see the data about column regarding NULL, defined as NO. Which means the field can't be null, we want it to be able to be null because it causes issues with later MySQL versions, which is why this update was made

Given the condition is true, the following runs

```text
ALTER TABLE `merc_stats` MODIFY COLUMN `special_abilities`  text CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL;
ALTER TABLE `npc_types` MODIFY COLUMN `special_abilities`  text CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL;
```

