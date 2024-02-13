# character_auras

## Relationships

```mermaid
erDiagram
    character_auras {
        int id
        int spell_id
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
    }
    spells_new {
        int id
        int descnum
        int effectdescnum
        int effectdescnum2
        int typedescnum
        varchar teleport_zone
    }
    character_auras ||--o{ character_data : "One-to-One"
    character_auras ||--o{ spells_new : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | id | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | spell_id | [spells_new](../../schema/spells/spells_new.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | [Unique Character Identifier](character_data.md) |
| slot | tinyint | Slot |
| spell_id | int | [Spell Identifier](../../schema/spells/spells_new.md) |

