# respawn_times

## Relationships

```mermaid
erDiagram
    respawn_times {
        smallint instance_id
        int id
    }
    spawn2 {
        varchar content_flags
        varchar content_flags_disabled
        int pathgrid
        int id
        int spawngroupID
        smallint version
        varchar zone
    }
    instance_list {
        int id
        tinyintunsigned version
        intunsigned zone
    }
    respawn_times ||--o{ spawn2 : "One-to-One"
    respawn_times ||--o{ instance_list : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | id | [spawn2](../../schema/spawns/spawn2.md) | id |
| One-to-One | instance_id | [instance_list](../../schema/instances/instance_list.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Respawn Time Identifier |
| start | int | Start UNIX Timestamp |
| duration | int | Duration in Seconds |
| instance_id | smallint | [Instance Identifier](../../schema/instances/instance_list.md) |

