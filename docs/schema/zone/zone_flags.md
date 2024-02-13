# zone_flags

## Relationships

```mermaid
erDiagram
    zone_flags {
        int charID
        int zoneID
        varchar zoneid
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
    zone_flags ||--o{ character_data : "One-to-One"
    zone_flags ||--o{ zone : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | charID | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | zoneid | [zone](../../schema/zone/zone.md) | zoneidnumber |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| charID | int | [Unique Character Identifier](../../schema/characters/character_data) |
| zoneID | int | [Zone Identifier](../../../../server/zones/zone-list) |

