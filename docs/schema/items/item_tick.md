# item_tick

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    item_tick {
        varchar it_itemid
        varchar it_qglobal
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
    quest_globals {
        int charid
        varchar name
        int npcid
        int zoneid
    }
    item_tick ||--o{ items : "One-to-One"
    item_tick ||--o{ quest_globals : "Has-Many"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | it_itemid | [items](../../schema/items/items.md) | id |
| Has-Many | it_qglobal | [quest_globals](../../schema/data-storage/quest_globals.md) | name |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| it_itemid | int | [Item Identifier](items.md) |
| it_chance | int | Chance: 0 = Never, 100 = Always |
| it_level | int | Level |
| it_id | int | [Spell Identifier](../../schema/spells/spells_new.md) |
| it_qglobal | varchar | [Quest Global Identifier](../../schema/data-storage/quest_globals.md) (Deprecated) |
| it_bagslot | tinyint | [Bag Slot](../../../../server/inventory/inventory-slots) |

