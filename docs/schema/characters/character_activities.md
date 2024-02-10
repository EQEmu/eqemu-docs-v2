# character_activities

## Relationships

```mermaid
erDiagram
    character_activities {
        intunsigned charid
        intunsigned taskid
        intunsigned activityid
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
    task_activities {
        intunsigned activityid
        varchar goalid
        varchar delivertonpc
        intunsigned taskid
        varchar zones
    }
    character_activities ||--o{ character_data : "One-to-One"
    character_activities ||--o{ tasks : "One-to-One"
    character_activities ||--o{ task_activities : "Has-Many"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | charid | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | taskid | [tasks](../../schema/tasks/tasks.md) | id |
| Has-Many | activityid | [task_activities](../../schema/tasks/task_activities.md) | activityid |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| charid | int | [Character Identifier](character_data.md) |
| taskid | int | [Task Identifier](../../schema/tasks/tasks.md) |
| activityid | int | [Activity Identifier](../../schema/tasks/task_activities.md) |
| donecount | int | Done Count |
| completed | tinyint | Completed: 0 = False, 1 = True |

