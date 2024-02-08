# completed_tasks

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    completed_tasks {
        intunsigned charid
        intunsigned taskid
        int activityid
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
    completed_tasks ||--o{ character_data : "One-to-One"
    completed_tasks ||--o{ tasks : "One-to-One"
    completed_tasks ||--o{ task_activities : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | charid | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | taskid | [tasks](../../schema/tasks/tasks.md) | id |
| One-to-One | activityid | [task_activities](../../schema/tasks/task_activities.md) | activityid |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| charid | int | [Character Identifier](../../schema/characters/character_data.md) |
| completedtime | int | Completed Time UNIX Timestamp |
| taskid | int | [Task Identifier](tasks.md) |
| activityid | int | [Activity Identifier](task_activities.md) |

