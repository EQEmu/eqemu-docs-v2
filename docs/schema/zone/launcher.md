# launcher

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    launcher {
        varchar name
    }
    launcher_zones {
        varchar launcher
        varchar zone
    }
    launcher ||--o{ launcher_zones : "Has-Many"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | name | [launcher_zones](../../schema/zone/launcher_zones.md) | launcher |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| name | varchar | Name |
| dynamics | tinyint | Dynamics |

