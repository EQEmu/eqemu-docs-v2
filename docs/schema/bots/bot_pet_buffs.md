# bot_pet_buffs

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    bot_pet_buffs {
        varchar pets_index
        varchar spell_id
    }
    bot_pets {
        varchar bot_id
        varchar pets_index
        varchar spell_id
    }
    spells_new {
        int id
        int descnum
        int effectdescnum
        int effectdescnum2
        int typedescnum
        varchar teleport_zone
    }
    bot_pet_buffs ||--o{ bot_pets : "One-to-One"
    bot_pet_buffs ||--o{ spells_new : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | pets_index | [bot_pets](../../schema/bots/bot_pets.md) | pets_index |
| One-to-One | spell_id | [spells_new](../../schema/spells/spells_new.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| pet_buffs_index | int | Unique Bot Pet Buffs Identifier |
| pets_index | int | [Bot Pet Identifier](bot_pets.md) |
| spell_id | int | [Spell Identifier](../../schema/spells/spells_new.md) |
| caster_level | int | Caster Level |
| duration | int | Duration of buff |

