# dynamic_zone_members

## Relationships

```mermaid
erDiagram
    dynamic_zone_members {
        intunsigned character_id
        intunsigned dynamic_zone_id
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
    }
    dynamic_zones {
        int dz_switch_id
        intunsigned id
        intunsigned compass_zone_id
        int instance_id
        intunsigned safe_return_zone_id
    }
    dynamic_zone_members ||--o{ character_data : "One-to-One"
    dynamic_zone_members ||--o{ dynamic_zones : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | character_id | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | dynamic_zone_id | [dynamic_zones](../../schema/tasks/shared_task_dynamic_zones.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Dynamic Zone Member Identifier |
| dynamic_zone_id | int | [Dynamic Zone Identifier](dynamic_zones.md) |
| character_id | int | [Character Identifier](../../schema/characters/character_data.md) |

