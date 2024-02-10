# character_disciplines

## Relationships

```mermaid
erDiagram
    character_disciplines {
        intunsigned id
        smallintunsigned disc_id
    }
    spells_new {
        int id
        int descnum
        int effectdescnum
        int effectdescnum2
        int typedescnum
        varchar teleport_zone
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
    }
    character_disciplines ||--o{ spells_new : "One-to-One"
    character_disciplines ||--o{ character_data : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | disc_id | [spells_new](../../schema/spells/spells_new.md) | id |
| One-to-One | id | [character_data](../../schema/characters/character_data.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | [Character Identifier](character_data.md) |
| slot_id | smallint | Slot Identifier |
| disc_id | smallint | [Discipline Identifier](../../schema/spells/spells_new.md) |

