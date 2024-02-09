# buyer

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    buyer {
        int charid
        int itemid
        varchar itemname
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
    buyer ||--o{ character_data : "One-to-One"
    buyer ||--o{ items : "One-to-One"
    buyer ||--o{ items : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | charid | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | itemid | [items](../../schema/items/items.md) | id |
| One-to-One | itemname | [items](../../schema/items/items.md) | name |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| charid | int | [Character Identifier](../../schema/characters/character_data.md) |
| buyslot | int | Slot |
| itemid | int | [Item Identifier](../../schema/items/items.md) |
| itemname | varchar | [Item Name](../../schema/items/items.md) |
| quantity | int | Quantity |
| price | int | Price |

