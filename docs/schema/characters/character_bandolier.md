# character_bandolier

## Relationships

```mermaid
erDiagram
    character_bandolier {
        intunsigned id
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
    character_bandolier ||--o{ character_data : "One-to-One"
    character_bandolier ||--o{ items : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | id | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | item_id | [items](../../schema/items/items.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | [Character Identifier](character_data.md) |
| bandolier_id | tinyint | Bandolier Identifier |
| bandolier_slot | tinyint | Slot |
| item_id | int | [Item Identifier](../../schema/items/items.md) |
| icon | int | Icon |
| bandolier_name | varchar | Name |

