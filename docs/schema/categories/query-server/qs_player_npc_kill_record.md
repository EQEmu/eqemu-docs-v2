# qs\_player\_npc\_kill\_record

| Column | Data Type | Description |
| :--- | :--- | :--- |
| fight\_id | int | Unique Fight Identifier |
| npc\_id | int | [NPC Type Identifier](https://github.com/EQEmu/docs-db-schema/tree/774e95edd473c84dafd6fe13b9b699f6b84a7ce8/docs/schema/categories/query_server/npc_types.md) |
| type | int | Type (2 - Raid, 1 - Group, 0 - Solo)|
| zone\_id | int | [Zone Identifier](https://eqemu.gitbook.io/server/categories/zones/zone-list) |
| time | timestamp | Time Timestamp |

