# character_task_timers

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    character_task_timers {
        intunsigned character_id
        intunsigned task_id
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
    }
    tasks {
        intunsigned id
        tinyint type
        intunsigned dz_template_id
    }
    character_task_timers ||--o{ character_data : "One-to-One"
    character_task_timers ||--o{ tasks : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | character_id | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | task_id | [tasks](../../schema/tasks/tasks.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Task Timer Identifier |
| character_id | int | [Character Identifer](character_data.md) |
| task_id | int | [Task Identifier](../../schema/tasks/tasks.md) |
| timer_type | int | Timer Type (0 = Replay, 1 = Request) |
| timer_group | int | Group Timer |
| expire_time | datetime | Expire Time |

