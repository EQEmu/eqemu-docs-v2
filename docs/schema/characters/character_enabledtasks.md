# character_enabledtasks

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    character_enabledtasks {
        intunsigned charid
        intunsigned taskid
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
    character_enabledtasks ||--o{ character_data : "One-to-One"
    character_enabledtasks ||--o{ tasks : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | charid | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | taskid | [tasks](../../schema/tasks/tasks.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| charid | int | [Character Identifier](character_data.md) |
| taskid | int | [Task Identifier](../../schema/tasks/tasks.md) |

