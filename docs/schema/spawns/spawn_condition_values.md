# spawn_condition_values

## Relationships

```mermaid
erDiagram
    spawn_condition_values {
        intunsigned instance_id
        intunsigned id
        varchar zone
    }
    spawn_conditions {
        mediumintunsigned id
        varchar zone
    }
    instance_list {
        int id
        tinyintunsigned version
        intunsigned zone
    }
    zone {
        int zoneidnumber
        varchar short_name
        tinyintunsigned version
        varchar content_flags
        varchar content_flags_disabled
    }
    spawn_condition_values ||--o{ spawn_conditions : "One-to-One"
    spawn_condition_values ||--o{ instance_list : "One-to-One"
    spawn_condition_values ||--o{ zone : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | id | [spawn_conditions](../../schema/spawns/spawn_conditions.md) | id |
| One-to-One | instance_id | [instance_list](../../schema/instances/instance_list.md) | id |
| One-to-One | zone | [zone](../../schema/zone/zone.md) | short_name |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | [Spawn Condition Identifier](spawn_conditions.md) |
| value | tinyint | Value |
| zone | varchar | [Zone Short Name](../../../../server/zones/zone-list) |
| instance_id | int | [Instance Identifier](../../schema/instances/instance_list.md) |

