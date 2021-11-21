# fishing

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Fishing Identifier |
| zoneid | int | [Zone Identifier](../../../../categories/zones/zone-list) |
| Itemid | int | [Item Identifier](../../../schema/categories/items/items.md) |
| skill_level | smallint | Skill Level |
| chance | smallint | Chance: 0 = Never, 100 = Always |
| npc_id | int | [NPC Type Identifier](../../../schema/categories/npcs/npc_types.md) |
| npc_chance | int | NPC Chance: 0 = Never, 100 = Always |
| min_expansion | tinyint |  |
| max_expansion | tinyint |  |
| content_flags | varchar |  |
| content_flags_disabled | varchar |  |

