# tasksets

## Relationships

```mermaid
erDiagram
    tasksets {
        intunsigned taskid
    }
    tasks {
        intunsigned id
        tinyint type
        intunsigned dz_template_id
    }
    tasksets ||--o{ tasks : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | taskid | [tasks](../../schema/tasks/tasks.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Task Set Entry Identifier |
| taskid | int | [Task Identifier](tasks.md) |

