# fishing

!!! info
	This page was last generated 2021.11.23

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Fishing Identifier |
| zoneid | int | [Zone Identifier](../../../../server/zones/zone-list) |
| Itemid | int | [Item Identifier](../../schema/items/items.md) |
| skill_level | smallint | Skill Level |
| chance | smallint | Chance: 0 = Never, 100 = Always |
| npc_id | int | [NPC Type Identifier](../../schema/npcs/npc_types.md) |
| npc_chance | int | NPC Chance: 0 = Never, 100 = Always |
| min_expansion | tinyint |  |
| max_expansion | tinyint |  |
| content_flags | varchar |  |
| content_flags_disabled | varchar |  |

