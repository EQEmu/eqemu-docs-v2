# npc\_faction\_entries

| Column | Data Type | Description |
| :--- | :--- | :--- |
| npc\_faction\_id | int | [NPC Faction Identifier](npc_faction.md) |
| faction\_id | int | [Faction Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/npcs/faction_list.md) |
| value | int | Value |
| npc\_value | tinyint | NPC Value: -1 = Attack, 0 = Neutral, 1 = Assist |
| temp | tinyint | Temeporary: 0 = Faction is permanent, player recieves a message, 1 = Faction is temporary, player does not recieve a message, 2 = Faction is temporary, player recieves a message, 3 = Faction is permanent, but player does not recieve a message. |

