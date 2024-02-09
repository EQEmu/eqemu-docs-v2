# starting_items

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    starting_items {
        varchar content_flags
        varchar content_flags_disabled
        varchar itemid
        varchar zone_id
        varchar zoneid
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
    starting_items ||--o{ content_flags : "One-to-One"
    starting_items ||--o{ content_flags : "One-to-One"
    starting_items ||--o{ items : "One-to-One"
    starting_items ||--o{ zone : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | content_flags | [content_flags](../../schema/flagging/content_flags.md) | flag_name |
| One-to-One | content_flags_disabled | [content_flags](../../schema/flagging/content_flags.md) | flag_name |
| One-to-One | itemid | [items](../../schema/items/items.md) | id |
| One-to-One | zone_id | [zone](../../schema/zone/zone.md) | zoneidnumber |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Starting Items Entry Identifier |
| race | int | [Race](../../../../server/npc/race-list): 0 = All |
| class | int | [Class](../../../../server/player/class-list): 0 = All |
| deity_list | text |  |
| zoneid | int | [Zone Identifier](../../../../server/zones/zone-list) |
| item_id | int |  |
| item_charges | tinyint | Item Charges |
| gm | tinyint | GM: 0 = False, 1 = True |
| inventory_slot | mediumint |  |
| min_expansion | tinyint | [Minimum Expansion](../../../../server/operation/expansion-list) |
| max_expansion | tinyint | [Maximum Expansion](../../../../server/operation/expansion-list) |
| content_flags | varchar | Content Flags Required to be Enabled |
| content_flags_disabled | varchar | Content Flags Required to be Disabled |

