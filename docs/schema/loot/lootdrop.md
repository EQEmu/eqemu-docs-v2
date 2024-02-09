# lootdrop

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    lootdrop {
        varchar content_flags
        varchar content_flags_disabled
        intunsigned id
    }
    content_flags {
        varchar flag_name
    }
    lootdrop_entries {
        int item_id
        intunsigned lootdrop_id
    }
    loottable_entries {
        intunsigned loottable_id
        intunsigned lootdrop_id
    }
    lootdrop ||--o{ content_flags : "One-to-One"
    lootdrop ||--o{ content_flags : "One-to-One"
    lootdrop ||--o{ lootdrop_entries : "Has-Many"
    lootdrop ||--o{ loottable_entries : "Has-Many"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | content_flags | [content_flags](../../schema/flagging/content_flags.md) | flag_name |
| One-to-One | content_flags_disabled | [content_flags](../../schema/flagging/content_flags.md) | flag_name |
| Has-Many | id | [lootdrop_entries](../../schema/loot/lootdrop_entries.md) | lootdrop_id |
| Has-Many | id | [loottable_entries](../../schema/loot/loottable_entries.md) | loottable_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Lootdrop Identifier |
| name | varchar | Name |
| min_expansion | tinyint | [Minimum Expansion](../../../../server/operation/expansion-list) |
| max_expansion | tinyint | [Maximum Expansion](../../../../server/operation/expansion-list) |
| content_flags | varchar | Content Flags Required to be Enabled |
| content_flags_disabled | varchar | Content Flags Required to be Disabled |

