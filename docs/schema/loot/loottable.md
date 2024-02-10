# loottable

## Relationships

```mermaid
erDiagram
    loottable {
        varchar content_flags
        varchar content_flags_disabled
        intunsigned id
    }
    content_flags {
        varchar flag_name
    }
    global_loot {
        varchar content_flags
        varchar content_flags_disabled
        int loottable_id
        mediumtext zone
    }
    loottable_entries {
        intunsigned loottable_id
        intunsigned lootdrop_id
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
    loottable ||--o{ content_flags : "One-to-One"
    loottable ||--o{ content_flags : "One-to-One"
    loottable ||--o{ global_loot : "Has-Many"
    loottable ||--o{ loottable_entries : "Has-Many"
    loottable ||--o{ npc_types : "Has-Many"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | content_flags | [content_flags](../../schema/flagging/content_flags.md) | flag_name |
| One-to-One | content_flags_disabled | [content_flags](../../schema/flagging/content_flags.md) | flag_name |
| Has-Many | id | [global_loot](../../schema/loot/global_loot.md) | loottable_id |
| Has-Many | id | [loottable_entries](../../schema/loot/loottable_entries.md) | loottable_id |
| Has-Many | id | [npc_types](../../schema/npcs/npc_types.md) | loottable_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Loottable Identifier |
| name | varchar | Name |
| mincash | int | Minimum Cash in Copper |
| maxcash | int | Maximum Cash in Copper |
| avgcoin | int | Average Coin in Copper |
| done | tinyint | Done: 0 = False, 1 = True |
| min_expansion | tinyint | [Minimum Expansion](../../../../server/operation/expansion-list) |
| max_expansion | tinyint | [Maximum Expansion](../../../../server/operation/expansion-list) |
| content_flags | varchar | Content Flags Required to be Enabled |
| content_flags_disabled | varchar | Content Flags Required to be Disabled |

