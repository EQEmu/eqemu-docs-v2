# instance_list

## Relationships

```mermaid
erDiagram
    instance_list {
        int id
        tinyintunsigned version
        intunsigned zone
    }
    adventure_details {
        intunsigned id
        smallintunsigned adventure_id
        int instance_id
    }
    character_bind {
        intunsigned id
        smallintunsigned zone_id
        mediumintunsigned instance_id
    }
    character_corpses {
        intunsigned id
        intunsigned charid
        varchar charname
        smallintunsigned instance_id
        smallint zone_id
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
    }
    character_instance_safereturns {
        intunsigned character_id
        int instance_zone_id
        int instance_id
        int safe_zone_id
    }
    instance_list ||--o{ adventure_details : "Has-Many"
    instance_list ||--o{ character_bind : "Has-Many"
    instance_list ||--o{ character_corpses : "Has-Many"
    instance_list ||--o{ character_data : "Has-Many"
    instance_list ||--o{ character_instance_safereturns : "Has-Many"
    instance_list ||--o{ character_instance_safereturns : "Has-Many"


```

```mermaid
erDiagram
    instance_list {
        int id
        tinyintunsigned version
        intunsigned zone
    }
    doors {
        varchar content_flags
        varchar content_flags_disabled
        int dz_switch_id
        int keyitem
        varchar zone
        varchar dest_zone
        intunsigned dest_instance
        smallint version
    }
    dynamic_zones {
        int dz_switch_id
        intunsigned id
        intunsigned compass_zone_id
        int instance_id
        intunsigned safe_return_zone_id
    }
    respawn_times {
        smallint instance_id
        int id
    }
    spawn_condition_values {
        intunsigned instance_id
        intunsigned id
        varchar zone
    }
    zone_points {
        varchar content_flags
        varchar content_flags_disabled
        intunsigned target_instance
        varchar zone
        intunsigned target_zone_id
        int version
    }
    instance_list ||--o{ doors : "Has-Many"
    instance_list ||--o{ dynamic_zones : "Has-Many"
    instance_list ||--o{ respawn_times : "Has-Many"
    instance_list ||--o{ spawn_condition_values : "Has-Many"
    instance_list ||--o{ zone_points : "Has-Many"


```

```mermaid
erDiagram
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
    instance_list ||--o{ zone : "One-to-One"
    instance_list ||--o{ zone : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [adventure_details](../../schema/adventures/adventure_details.md) | instance_id |
| Has-Many | id | [character_bind](../../schema/characters/character_bind.md) | instance_id |
| Has-Many | id | [character_corpses](../../schema/characters/character_corpses.md) | instance_id |
| Has-Many | id | [character_data](../../schema/characters/character_data.md) | zone_instance |
| Has-Many | id | [character_instance_safereturns](../../schema/characters/character_instance_safereturns.md) | instance_id |
| Has-Many | id | [character_instance_safereturns](../../schema/characters/character_instance_safereturns.md) | instance_zone_id |
| Has-Many | id | [doors](../../schema/doors/doors.md) | dest_instance |
| Has-Many | id | [dynamic_zones](../../schema/tasks/shared_task_dynamic_zones.md) | instance_id |
| Has-Many | id | [respawn_times](../../schema/spawns/respawn_times.md) | instance_id |
| Has-Many | id | [spawn_condition_values](../../schema/spawns/spawn_condition_values.md) | instance_id |
| Has-Many | id | [zone_points](../../schema/zone/zone_points.md) | target_instance |
| One-to-One | version | [zone](../../schema/zone/zone.md) | version |
| One-to-One | zone | [zone](../../schema/zone/zone.md) | zoneidnumber |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Instance Identifier |
| zone | int | [Zone Identifier](../../../../server/zones/zone-list) |
| version | tinyint | Version |
| is_global | tinyint | Is Global: 0 = False, 1 = True |
| start_time | int | Start Time UNIX Timestamp |
| duration | int | Duration in Seconds |
| never_expires | tinyint | Never Expires: 0 = False, 1 = True |
| notes | varchar |  |

