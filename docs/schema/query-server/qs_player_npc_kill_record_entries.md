# qs_player_npc_kill_record_entries

## Relationships

```mermaid
erDiagram
    qs_player_npc_kill_record_entries {
        int char_id
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
    }
    qs_player_npc_kill_record_entries ||--o{ character_data : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | char_id | [character_data](../../schema/characters/character_data.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| event_id | int | Unique Event Identifier |
| char_id | int | [Character Identifier](../../schema/characters/character_data.md) |

