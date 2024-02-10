# qs_player_move_record

## Relationships

```mermaid
erDiagram
    qs_player_move_record {
        int char_id
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
    }
    qs_player_move_record ||--o{ character_data : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | char_id | [character_data](../../schema/characters/character_data.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| move_id | int | Unique Move Identifier |
| time | timestamp | Time Timestamp |
| char_id | int | [Character Identifier](../../schema/characters/character_data.md) |
| from_slot | mediumint | [From Slot Identifier](../../../../server/inventory/inventory-slots) |
| to_slot | mediumint | [To Slot Identifier](../../../../server/inventory/inventory-slots) |
| stack_size | mediumint | Stack Size |
| char_items | mediumint | [Character Item Identifier](../../schema/items/items.md) |
| postaction | tinyint | Post Action |

