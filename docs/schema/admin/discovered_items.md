# discovered_items

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    discovered_items {
        varchar char_name
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
    discovered_items ||--o{ character_data : "One-to-One"
    discovered_items ||--o{ items : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | char_name | [character_data](../../schema/characters/character_data.md) | name |
| One-to-One | item_id | [items](../../schema/items/items.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| item_id | int | [Item Identifier](../../schema/items/items.md) |
| char_name | varchar | [Character Name](../../schema/characters/character_data.md) |
| discovered_date | int | Discovered Date UNIX Timestamp |
| account_status | int | [Account Status](../../../../server/player/status-levels) |

