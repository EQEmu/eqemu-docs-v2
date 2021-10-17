# merc\_armorinfo

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Mercenary Armor Info Identifier |
| merc\_npc\_type\_id | int | [Mercenary NPC Type Identifier](merc_npc_types.md) |
| minlevel | tinyint | Minimum Level |
| maxlevel | tinyint | Maximum Level |
| texture | tinyint | [Texture](https://eqemu.gitbook.io/server/categories/npc/textures) |
| helmtexture | tinyint | [Helmet Texture](https://eqemu.gitbook.io/server/categories/npc/textures) |
| armortint\_id | int | [Armor Tint Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/mercenaries/npc_types_tint.md) |
| armortint\_red | tinyint | Armor Tint Red: 0 = None, 255 = Max |
| armortint\_green | tinyint | Armor Tint Green: 0 = None, 255 = Max |
| armortint\_blue | tinyint | Armor Tint Blue: 0 = None, 255 = Max |

