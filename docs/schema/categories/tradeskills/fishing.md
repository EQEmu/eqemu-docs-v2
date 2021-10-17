# fishing

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Fishing Identifier |
| zoneid | int | [Zone Identifier](https://eqemu.gitbook.io/server/categories/zones/zone-list) |
| Itemid | int | [Item Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/tradeskills/items.md) |
| skill\_level | smallint | Skill Level |
| chance | smallint | Chance: 0 = Never, 100 = Always |
| npc\_id | int | [NPC Type Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/tradeskills/npc_types.md) |
| npc\_chance | int | NPC Chance: 0 = Never, 100 = Always |

