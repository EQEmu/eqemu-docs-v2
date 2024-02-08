# spawn_conditions

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    spawn_conditions {
        mediumintunsigned id
        varchar zone
    }
    spawn_condition_values {
        intunsigned instance_id
        intunsigned id
        varchar zone
    }
    zone {
        int zoneidnumber
        varchar short_name
        tinyintunsigned version
        varchar content_flags
        varchar content_flags_disabled
    }
    spawn_conditions ||--o{ spawn_condition_values : "One-to-One"
    spawn_conditions ||--o{ zone : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | id | [spawn_condition_values](../../schema/spawns/spawn_condition_values.md) | id |
| One-to-One | zone | [zone](../../schema/zone/zone.md) | short_name |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| zone | varchar | [Zone Short Name](../../../../server/zones/zone-list) |
| id | mediumint | Spawn Condition Identifier |
| value | mediumint | Value |
| onchange | tinyint | [On Change Type](../../../../server/npc/spawns/on-change-types) |
| name | varchar | Name |

