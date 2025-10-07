---
description: This article describes the ability of versioning and patching your database changes
---

Version patching is the idea of taking a base PEQ database and "patching it" with changes for your custom server, in a way it can be ran multiple times, and versioned in git.

This is a discipline that lends itself to custom server changes that you want to be destructive and have an "undo button" with (reset database to base, and re-inject changes).

All links are referenced to an archive server called rebuildeq/server

## Makefile Injects

In this [Makefile](https://github.com/rebuildeq/server/blob/main/.devcontainer/Makefile#L335) you can see how a new `make inject-custom` was added.

- The first line has `@cd base && make inject --no-print-directory`, this went to the [base](https://github.com/rebuildeq/server/blob/main/base/Makefile) subfolder and ran the `make inject` command. Each inject was categorized by a subfolder. We'll use item as an example.

- Inside [base/item](https://github.com/rebuildeq/server/tree/main/base/item) subfolder, there is a [Makefile](<https://github.com/rebuildeq/server/blob/main/base/item/Makefile>) that sourced all the sql files in this directory (salvaging-skills.sql, boxes.sql, parts.sql, etc) and any subfolder sql files.

- These sql files are all built to be applied multiple times with no side effects. Such as:

- [aug.sql](https://github.com/rebuildeq/server/blob/main/base/item/aug/aug.sql) patched aug slots to a new simplified variant.


## card.sql

[card.sql](https://github.com/rebuildeq/server/blob/main/base/item/aug/card.sql) is a custom aug system introduced by the server,


The first line (REPLACE INTO) is what I call a "template" item, this was inserted into the items table, and was the baseline all cards derived from.


Temporary table [line 6](<https://github.com/rebuildeq/server/blob/main/base/item/aug/card.sql#L6>), this is a temporary table labeled `card_base_table`. What this does is it creates a temporary table that takes the first record of the template, and puts it into a temporary table, so it can be manipulated quickly, and then placed back in the normal items table once applied.


Stored Procedure `tt_start` [line 9-13](<https://github.com/rebuildeq/server/blob/main/base/item/aug/card.sql#L9>), this takes the temporary card_base_table and copies it to another temporary table item_table, this is where you can manipulate the template with your changes.

Stored Procedure `tt_end` [line 15-20(<https://github.com/rebuildeq/server/blob/main/base/item/aug/card.sql#L15>), this takes the temporary item_table and does an REPLACE INTO into the real items table, this means if the item already exists, it will be updated, if it doesn't exist, it will be inserted.

Some variables get SET on [lines 24-27](<https://github.com/rebuildeq/server/blob/main/base/item/aug/card.sql#L24>):
```sql
SET @Chest = 131072;
SET @Legs = 262144;
SET @Finger = 98304;
SET @Weapon = 8192;
```
These can be used during injects so you don't have to hard code numbers and keep it a bit more human readable.

Now on to the real intent: [inserting data](<https://github.com/rebuildeq/server/blob/main/base/item/aug/card.sql#L30C1-L32C134>)
```sql
CALL tt_start; UPDATE item_table SET id=148101, name='Dragon Card', 			slots=@Chest, icon=6152, manaregen=2 ; CALL tt_end;
CALL tt_start; UPDATE item_table SET id=148102, name='Insect Card', 			slots=@Chest, icon=6170, hp=10, ac=10; CALL tt_end;
CALL tt_start; UPDATE item_table SET id=148103, name='Animal Card', 			slots=@Chest, icon=6145, ac=10, aagi=5, attack=5; CALL tt_end;
```

This is what ultimately does the changes.
- Call tt_start, to create a copy of the template in a temporary table
- Do edits to the template unique for the item
- Call tt_end, to apply the changes to the real items table

Now a single line can do all these operations, making it easy to read, and easy to version control, and edit in the future.

