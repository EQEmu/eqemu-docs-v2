# npc_faction

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    npc_faction {
        int id
        int primaryfaction
    }
    npc_faction_entries {
        intunsigned npc_faction_id
        intunsigned faction_id
    }
    faction_list {
        int id
    }
    npc_faction ||--o{ npc_faction_entries : "Has-Many"
    npc_faction ||--o{ faction_list : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [npc_faction_entries](../../schema/npcs/npc_faction_entries.md) | npc_faction_id |
| One-to-One | primaryfaction | [faction_list](../../schema/factions/faction_list.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique NPC Faction Identifier |
| name | tinytext | Name |
| primaryfaction | int | [Primary Faction Identifier](../../schema/factions/faction_list.md) |
| ignore_primary_assist | tinyint | Ignore Primary Assist: 0 = False, &gt;0 = True |

