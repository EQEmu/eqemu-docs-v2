# npc_spells_effects

## Relationships

```mermaid
erDiagram
    npc_spells_effects {
        intunsigned id
    }
    npc_spells_effects_entries {
        int npc_spells_effects_id
    }
    npc_spells_effects ||--o{ npc_spells_effects_entries : "Has-Many"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [npc_spells_effects_entries](../../schema/npcs/npc_spells_effects_entries.md) | npc_spells_effects_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Spell Effects Identifier |
| name | tinytext | Name |
| parent_list | int | Parent List |

