# character_pet_info

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    character_pet_info {
        int char_id
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
    character_pet_info ||--o{ character_data : "One-to-One"
    character_pet_info ||--o{ spells_new : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | char_id | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | spell_id | [spells_new](../../schema/spells/spells_new.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| char_id | int | [Character Identifier](character_data.md) |
| pet | int | Pet |
| petname | varchar | Pet Name |
| petpower | int | Pet Power |
| spell_id | int | [Spell Identifier](../../schema/spells/spells_new.md) |
| hp | int | Health |
| mana | int | Mana |
| size | float | Size |
| taunting | tinyint | Taunting: 0 = False, 1 = True |

