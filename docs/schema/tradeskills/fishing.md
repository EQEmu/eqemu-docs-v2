# fishing

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    fishing {
        varchar content_flags
        varchar content_flags_disabled
        int Itemid
        int zoneid
        int npc_id
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
    npc_types {
        intunsigned alt_currency_id
        int id
        int npc_faction_id
        text name
        intunsigned loottable_id
        intunsigned merchant_id
        intunsigned emoteid
        intunsigned adventure_template_id
        intunsigned armortint_id
        intunsigned npc_spells_id
        intunsigned npc_spells_effects_id
        intunsigned trap_template
    }
    fishing ||--o{ content_flags : "One-to-One"
    fishing ||--o{ content_flags : "One-to-One"
    fishing ||--o{ items : "One-to-One"
    fishing ||--o{ zone : "One-to-One"
    fishing ||--o{ npc_types : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | content_flags | [content_flags](../../schema/flagging/content_flags.md) | flag_name |
| One-to-One | content_flags_disabled | [content_flags](../../schema/flagging/content_flags.md) | flag_name |
| One-to-One | Itemid | [items](../../schema/items/items.md) | id |
| One-to-One | zoneid | [zone](../../schema/zone/zone.md) | zoneidnumber |
| One-to-One | npc_id | [npc_types](../../schema/npcs/npc_types.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Fishing Identifier |
| zoneid | int | [Zone Identifier](../../../../server/zones/zone-list) |
| Itemid | int | [Item Identifier](../../schema/items/items.md) |
| skill_level | smallint | Skill Level |
| chance | smallint | Chance: 0 = Never, 100 = Always |
| npc_id | int | [NPC Type Identifier](../../schema/npcs/npc_types.md) |
| npc_chance | int | NPC Chance: 0 = Never, 100 = Always |
| min_expansion | tinyint | [Minimum Expansion](../../../../server/operation/expansion-list) |
| max_expansion | tinyint | [Maximum Expansion](../../../../server/operation/expansion-list) |
| content_flags | varchar | Content Flags Required to be Enabled |
| content_flags_disabled | varchar | Content Flags Required to be Disabled |

