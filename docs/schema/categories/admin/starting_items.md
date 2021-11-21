# starting_items

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Starting Items Entry Identifier |
| race | int | [Race](../../../../categories/npc/race-list): 0 = All |
| class | int | [Class](../../../../categories/player/class-list): 0 = All |
| deityid | int | [Deity](../../../../categories/player/deity-list): 0 = Alll |
| zoneid | int | [Zone Identifier](../../../../categories/zones/zone-list) |
| itemid | int | [Item Identifier](../../../schema/categories/items/items.md) |
| item_charges | tinyint | Item Charges |
| gm | tinyint | GM: 0 = False, 1 = True |
| slot | mediumint | [Slot](../../../../categories/inventory/inventory-slots) |

