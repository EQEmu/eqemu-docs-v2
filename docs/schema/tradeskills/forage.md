# forage

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    forage {
        varchar content_flags
        varchar content_flags_disabled
        int Itemid
        int zoneid
    }
    content_flags {
        varchar flag_name
    }
    items {
        int id
        int book
        varchar name
        int recasttype
        int icon
        mediumint bardeffect
        int clickeffect
        int focuseffect
        int proceffect
        int scrolleffect
        int worneffect
    }
    zone {
        int zoneidnumber
        varchar short_name
        tinyintunsigned version
        varchar content_flags
        varchar content_flags_disabled
    }
    forage ||--o{ content_flags : "One-to-One"
    forage ||--o{ content_flags : "One-to-One"
    forage ||--o{ items : "One-to-One"
    forage ||--o{ zone : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | content_flags | [content_flags](../../schema/flagging/content_flags.md) | flag_name |
| One-to-One | content_flags_disabled | [content_flags](../../schema/flagging/content_flags.md) | flag_name |
| One-to-One | Itemid | [items](../../schema/items/items.md) | id |
| One-to-One | zoneid | [zone](../../schema/zone/zone.md) | zoneidnumber |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Forage Identifier |
| zoneid | int | [Zone Identifier](../../../../server/zones/zone-list) |
| Itemid | int | [Item Identifier](../../schema/items/items.md) |
| level | smallint | Level |
| chance | smallint | Chance: 0 = Never, 100 = Always |
| min_expansion | tinyint | [Minimum Expansion](../../../../server/operation/expansion-list) |
| max_expansion | tinyint | [Maximum Expansion](../../../../server/operation/expansion-list) |
| content_flags | varchar | Content Flags Required to be Enabled |
| content_flags_disabled | varchar | Content Flags Required to be Disabled |

