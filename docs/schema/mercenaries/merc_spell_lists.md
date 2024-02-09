# merc_spell_lists

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    merc_spell_lists {
        varchar merc_spell_list_id
    }
    merc_spell_list_entries {
        varchar merc_spell_list_id
        varchar spell_id
        varchar stance_id
    }
    merc_spell_lists ||--o{ merc_spell_list_entries : "Has-Many"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | merc_spell_list_id | [merc_spell_list_entries](../../schema/mercenaries/merc_spell_list_entries.md) | merc_spell_list_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| merc_spell_list_id | int | Unique Mercenary Spell List Identifier |
| class_id | int | [Class Identifier](../../../../server/player/class-list) |
| proficiency_id | tinyint | Proficiency Identifier |
| name | varchar | Name |

