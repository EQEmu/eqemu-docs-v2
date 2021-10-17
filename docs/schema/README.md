---
description: >-
  This repository is purely used to sync database documentation with
  https://www.gitbook.com/
---

# Introduction

## Database Conventions and Guidelines

This page serves as a reference for rules that we adhere to as a project, things could change over time but this is mostly a living representation of our current spec.

This does not mean that past table creations match this specification, but that we intend to keep it consistent going forward.

## Contributing

For anyone interested in contributing to the database documentation.

If a database table's page has an inaccuracy in a column, modify the database-schema-reference.yml.

This is where all the table and column information is stored.

Please do not submit pull requests modifying the .md files directly, as they will be over-written by the doc-gen.js.

## Table Names

* lowercase
* snake\_case
* Plural

Table names should be **lowercase**, `snake_case` and should clearly describe the purpose of the table itself.

Table names should also adhere to an appropriate category prefix if necessary. For example, if the table is storing data that is particular to that of a character, it is appropriate to prefix the table name with `character_x`

**Examples**

```text
character_auras
character_bandolier
character_buffs
character_disciplines
```

When defining an object or model in the code, our source is inconsistent everywhere, but trend to use the singular name of an entity as an object and the table name as plural.

For example I have a class representing a `Door` and a table named `doors`

## Column Naming

Like **Table Names** columns also adhere to **lowercase** and `snake_case` appropriations. The column itself should very easily describe the purpose of the column itself without abbreviations as much as possible.

For example, instead of `p_cp`, it is far easier for new server operators and developers to understand `player_copper`, don't be lazy and don't be afraid to be verbose.

### Foreign Key Consistency

If your column has a relationship to another table, make sure that it prefixes the table name with id.

**Example**

I'm making a new table called `keyring` and the schema looks like this:

```text
id - int(11) pri - key
character_id - int(11)
item_id - int(11)
```

We easily know that we have a loose foreign key relationship to the `character_data` table \(Which currently breaks convention and should be called `characters`\).

We also know that we have a loose relationship to the `items` table and we resolve to `items:id`

## Have an Integer Primary Key

At minimum, add the standard `id` column with an auto-incrementing integer. This makes sequencing easier.

## Store Datetimes as Datetimes

We can easily convert to and from unix using datetime, use this as a standard practice.

## Indexes

A simple index can go a long way for performance if you have data that is being looked up frequently especially in the case of strings.

For example, we have a table called `saylink` \(should be plural\) that contains `phrase` which gets looked up frequently when a saylink is clicked or when saylinks are being parsed inside of a `quest::say` context, this lookup and scan gets expensive when there is no index on the column itself. What ends up happening is that the MySQL engine ends up having to do full table scans to find the phrase corresponding to the requested record to see if it exists or lookup and ID associated to said saylink.

If your table's primary method of lookup is through `id` - you already get indexing out of the box, there is no additional indexes required.

## Unsigned Versus Signed

Only use what you intend to use for your integer space, if you don't plan on having negative values, make your field unsigned and corresponding C/C++ datatype to match. Or, use a bigger data integer type as signed to store your unsigned value.

## Soft Deletes

If your table or feature uses the concept of soft deleting an object, please use `deleted_at` in a `datetime` field to mark that entity as deleted and then make sure you use queries that take into consideration where `deleted_at` is null \(An index may be appropriate on this field\).

