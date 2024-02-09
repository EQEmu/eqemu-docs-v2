# horses

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    horses {
        varchar filename
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
    horses ||--o{ npc_types : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | filename | [npc_types](../../schema/npcs/npc_types.md) | name |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Horse Identifier |
| filename | varchar | [NPC Type Name](../../schema/npcs/npc_types.md) |
| race | smallint | [Race](../../../../server/npc/race-list) |
| gender | tinyint | [Gender](../../../../server/npc/genders) |
| texture | tinyint | [Texture](../../../../server/npc/textures) |
| mountspeed | float | Mount Run Speed |
| notes | varchar | Notes |

