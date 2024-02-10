# shared_tasks

## Relationships

```mermaid
erDiagram
    shared_tasks {
        bigint id
        int task_id
    }
    shared_task_activity_state {
        int activity_id
        bigint shared_task_id
    }
    shared_task_dynamic_zones {
        intunsigned dynamic_zone_id
        bigint shared_task_id
    }
    shared_task_members {
        bigint character_id
        bigint shared_task_id
    }
    tasks {
        intunsigned id
        tinyint type
        intunsigned dz_template_id
    }
    shared_tasks ||--o{ shared_task_activity_state : "Has-Many"
    shared_tasks ||--o{ shared_task_dynamic_zones : "Has-Many"
    shared_tasks ||--o{ shared_task_members : "Has-Many"
    shared_tasks ||--o{ tasks : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [shared_task_activity_state](../../schema/tasks/shared_task_activity_state.md) | shared_task_id |
| Has-Many | id | [shared_task_dynamic_zones](../../schema/tasks/shared_task_dynamic_zones.md) | shared_task_id |
| Has-Many | id | [shared_task_members](../../schema/tasks/shared_task_members.md) | shared_task_id |
| One-to-One | task_id | [tasks](../../schema/tasks/tasks.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | bigint | Unique Shared Task Identifier |
| task_id | int | [Task Identifier](tasks.md) |
| accepted_time | datetime | Accepted Time |
| expire_time | datetime | Expire Time |
| completion_time | datetime | Completion Time |
| is_locked | tinyint | Is Locked: 0 = False, 1 = True |

