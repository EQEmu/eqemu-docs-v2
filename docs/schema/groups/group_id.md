# group_id

## Relationships

```mermaid
erDiagram
    group_id {
        int charid
        int groupid
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
    }
    group_leaders {
        varchar assist
        varchar leadername
        varchar maintank
        varchar mentoree
        varchar puller
        int gid
    }
    group_id ||--o{ character_data : "One-to-One"
    group_id ||--o{ group_leaders : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | charid | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | groupid | [group_leaders](../../schema/groups/group_leaders.md) | gid |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| groupid | int | Unique Group Identifier |
| charid | int | [Character Identifier](../../schema/characters/character_data.md) |
| name | varchar | Name |
| ismerc | tinyint | Is Mercenary: 0 = False, 1 = True |

