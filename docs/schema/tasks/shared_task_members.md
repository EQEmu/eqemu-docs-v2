# shared_task_members

## Relationships

```mermaid
erDiagram
    shared_task_members {
        bigint character_id
        bigint shared_task_id
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
    }
    shared_tasks {
        bigint id
        int task_id
    }
    shared_task_members ||--o{ character_data : "One-to-One"
    shared_task_members ||--o{ shared_tasks : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | character_id | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | shared_task_id | [shared_tasks](../../schema/tasks/shared_tasks.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| shared_task_id | bigint | [Shared Task Identifier](shared_tasks.md) |
| character_id | bigint | [Character Identifier](../../schema/characters/character_data.md) |
| is_leader | tinyint | Is Leader: 0 = False, 1 = True |

