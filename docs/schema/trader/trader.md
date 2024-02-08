# trader

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    trader {
        intunsigned char_id
        intunsigned item_id
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
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
    trader ||--o{ character_data : "One-to-One"
    trader ||--o{ items : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | char_id | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | item_id | [items](../../schema/items/items.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| char_id | int | [Unique Character Identifier](../../schema/characters/character_data.md) |
| item_id | int | [Item Identifier](../../schema/items/items.md) |
| serialnumber | int | Serial Number |
| charges | int | Charges |
| item_cost | int | Item Cost |
| slot_id | tinyint | [Slot Identifier](../../../../server/inventory/inventory-slots) |

