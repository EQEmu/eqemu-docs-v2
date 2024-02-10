# ldon_trap_templates

## Relationships

```mermaid
erDiagram
    ldon_trap_templates {
        intunsigned id
        smallintunsigned spell_id
    }
    ldon_trap_entries {
        intunsigned trap_id
    }
    spells_new {
        int id
        int descnum
        int effectdescnum
        int effectdescnum2
        int typedescnum
        varchar teleport_zone
    }
    ldon_trap_templates ||--o{ ldon_trap_entries : "Has-Many"
    ldon_trap_templates ||--o{ spells_new : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [ldon_trap_entries](../../schema/traps/ldon_trap_entries.md) | trap_id |
| One-to-One | spell_id | [spells_new](../../schema/spells/spells_new.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique LDoN Trap Template Identifier |
| type | tinyint | [Trap Type](../../../../server/zones/trap-types) |
| spell_id | smallint | [Spell Identifier](../../schema/spells/spells_new.md) |
| skill | smallint | [Skill](../../../../server/player/skills) |
| locked | tinyint | Locked: 0 = False, 1 = True |

