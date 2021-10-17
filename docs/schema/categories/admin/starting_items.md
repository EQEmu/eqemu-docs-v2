# starting\_items

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Starting Items Entry Identifier |
| race | int | [Race](https://eqemu.gitbook.io/server/categories/npc/race-list): 0 = All |
| class | int | [Class](https://eqemu.gitbook.io/server/categories/player/class-list): 0 = All |
| deityid | int | [Deity](https://eqemu.gitbook.io/server/categories/player/deity-list): 0 = Alll |
| zoneid | int | [Zone Identifier](https://eqemu.gitbook.io/server/categories/zones/zone-list) |
| itemid | int | [Item Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/admin/items.md) |
| item\_charges | tinyint | Item Charges |
| gm | tinyint | GM: 0 = False, 1 = True |
| slot | mediumint | [Slot](https://eqemu.gitbook.io/server/categories/inventory/inventory-slots) |

