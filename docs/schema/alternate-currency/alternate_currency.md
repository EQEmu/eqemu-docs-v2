# alternate_currency

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    alternate_currency {
        int id
        int item_id
    }
    character_alt_currency {
        intunsigned currency_id
        intunsigned char_id
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
    alternate_currency ||--o{ character_alt_currency : "Has-Many"
    alternate_currency ||--o{ npc_types : "Has-Many"
    alternate_currency ||--o{ items : "Has-Many"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [character_alt_currency](../../schema/characters/character_alt_currency.md) | currency_id |
| Has-Many | id | [npc_types](../../schema/npcs/npc_types.md) | alt_currency_id |
| Has-Many | item_id | [items](../../schema/items/items.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | [Alternate Currency Identifier](../../../../server/items/alternate-currencies) |
| item_id | int | [Item Identifier](../../schema/items/items.md) |

