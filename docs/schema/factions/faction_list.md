# faction_list

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    faction_list {
        int id
    }
    faction_list_mod {
        intunsigned faction_id
    }
    npc_types {
        intunsigned alt_currency_id
        int id
        int npc_faction_id
        text name
        intunsigned loottable_id
        intunsigned merchant_id
        intunsigned emoteid
        intunsigned adventure_template_id
        intunsigned armortint_id
        intunsigned npc_spells_id
        intunsigned npc_spells_effects_id
        intunsigned trap_template
    }
    faction_list ||--o{ faction_list_mod : "Has-Many"
    faction_list ||--o{ npc_types : "Has-Many"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [faction_list_mod](../../schema/factions/faction_list_mod.md) | faction_id |
| Has-Many | id | [npc_types](../../schema/npcs/npc_types.md) | npc_faction_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Faction Identifier |
| name | varchar | Name |
| base | smallint | Base |

