# character_bind

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    character_bind {
        intunsigned id
        smallintunsigned zone_id
        mediumintunsigned instance_id
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
    }
    zone {
        int zoneidnumber
        varchar short_name
        tinyintunsigned version
        varchar content_flags
        varchar content_flags_disabled
    }
    instance_list {
        int id
        tinyintunsigned version
        intunsigned zone
    }
    character_bind ||--o{ character_data : "One-to-One"
    character_bind ||--o{ zone : "One-to-One"
    character_bind ||--o{ instance_list : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | id | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | zone_id | [zone](../../schema/zone/zone.md) | zoneidnumber |
| One-to-One | instance_id | [instance_list](../../schema/instances/instance_list.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | [Character Identifier](character_data.md) |
| slot | int | Slot |
| zone_id | smallint | [Zone Identifier](../../../../server/zones/zone-list) |
| instance_id | mediumint | Instance Identifier |
| x | float | X Coordinate |
| y | float | Y Coordinate |
| z | float | Z Coordinate |
| heading | float | Heading Coordinate |

