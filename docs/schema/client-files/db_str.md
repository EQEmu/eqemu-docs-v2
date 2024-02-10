# db_str

## Relationships

```mermaid
erDiagram
    db_str {
        int id
    }
    merc_templates {
        varchar dbstring
        varchar merc_template_id
        varchar name_type_id
        varchar merc_npc_type_id
        varchar merc_subtype_id
        varchar merc_type_id
    }
    merc_types {
        varchar dbstring
        varchar merc_type_id
    }
    spells_new {
        int id
        int descnum
        int effectdescnum
        int effectdescnum2
        int typedescnum
        varchar teleport_zone
    }
    db_str ||--o{ merc_templates : "One-to-One"
    db_str ||--o{ merc_types : "One-to-One"
    db_str ||--o{ spells_new : "One-to-One"
    db_str ||--o{ spells_new : "One-to-One"
    db_str ||--o{ spells_new : "One-to-One"
    db_str ||--o{ spells_new : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | id | [merc_templates](../../schema/mercenaries/merc_templates.md) | dbstring |
| One-to-One | id | [merc_types](../../schema/mercenaries/merc_types.md) | dbstring |
| One-to-One | id | [spells_new](../../schema/spells/spells_new.md) | descnum |
| One-to-One | id | [spells_new](../../schema/spells/spells_new.md) | effectdescnum |
| One-to-One | id | [spells_new](../../schema/spells/spells_new.md) | effectdescnum2 |
| One-to-One | id | [spells_new](../../schema/spells/spells_new.md) | typedescnum |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Database String Identifier |
| type | int | Type |
| value | text | Value |

