# faction_base_data

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    faction_base_data {
        smallint client_faction_id
    }
    faction_list {
        int id
    }
    faction_base_data ||--o{ faction_list : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | client_faction_id | [faction_list](../../schema/factions/faction_list.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| client_faction_id | smallint | [Client Faction Identifier](faction_list.md) |
| min | smallint | Minimum Faction |
| max | smallint | Maximum Faction |
| unk_hero1 | smallint | Unknown |
| unk_hero2 | smallint | Unknown |
| unk_hero3 | smallint | Unknown |

