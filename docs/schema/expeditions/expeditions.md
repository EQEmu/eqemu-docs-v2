# expeditions

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    expeditions {
        intunsigned id
        intunsigned dynamic_zone_id
    }
    dynamic_zones {
        int dz_switch_id
        intunsigned id
        intunsigned compass_zone_id
        int instance_id
        intunsigned safe_return_zone_id
    }
    expeditions ||--o{ dynamic_zones : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | dynamic_zone_id | [dynamic_zones](../../schema/tasks/shared_task_dynamic_zones.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Expedition Identifier |
| dynamic_zone_id | int | [Dynamic Zone Identifier](../../schema/dynamic-zones/dynamic_zones.md) |
| add_replay_on_join | tinyint | Add Replay On Join: 0 = False, 1 = True |
| is_locked | tinyint | Is Locked: 0 = False, 1 = True |

