# lootdrop_entries

## Relationships

```mermaid
erDiagram
    lootdrop_entries {
        int item_id
        intunsigned lootdrop_id
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
    lootdrop {
        varchar content_flags
        varchar content_flags_disabled
        intunsigned id
    }
    lootdrop_entries ||--o{ items : "One-to-One"
    lootdrop_entries ||--o{ lootdrop : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | item_id | [items](../../schema/items/items.md) | id |
| One-to-One | lootdrop_id | [lootdrop](../../schema/loot/lootdrop.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| lootdrop_id | int | [Lootdrop Identifier](lootdrop.md) |
| item_id | int | [Item Identifier](../../schema/items/items.md) |
| item_charges | smallint | Item Charges |
| equip_item | tinyint | Equip Item: 0 = False, 1 = True |
| chance | float | Chance: 0 = Never, 100 = Always |
| disabled_chance | float | Disabled Chance: 0 = Never, 100 = Always |
| trivial_min_level | smallint | Trivial Minimum Level |
| trivial_max_level | smallint | Trivial Maximum Level |
| multiplier | tinyint | Multiplier |
| npc_min_level | smallint | NPC Minimum Level |
| npc_max_level | smallint | NPC Maximum Level |

