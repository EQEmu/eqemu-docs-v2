# character_tasks

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    character_tasks {
        intunsigned charid
        intunsigned taskid
        tinyint type
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
    character_tasks ||--o{ character_data : "One-to-One"
    character_tasks ||--o{ tasks : "One-to-One"
    character_tasks ||--o{ tasks : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | charid | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | taskid | [tasks](../../schema/tasks/tasks.md) | id |
| One-to-One | type | [tasks](../../schema/tasks/tasks.md) | type |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| charid | int | [Character Identifier](character_data.md) |
| taskid | int | [Task Identifier](../../schema/tasks/tasks.md) |
| slot | int | Slot |
| type | tinyint | [Task Type](../../../../server/task-system-guide/task-types) |
| acceptedtime | int | Accepted Time UNIX Timestamp |
| was_rewarded | tinyint | Was Rewarded: 0 = False, 1 = True |

