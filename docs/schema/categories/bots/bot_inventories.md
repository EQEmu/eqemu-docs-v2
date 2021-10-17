# bot\_inventories

| Column | Data Type | Description |
| :--- | :--- | :--- |
| inventories\_index | int | Unique Bot Inventory Identifier |
| bot\_id | int | [Bot Identifier](bot_data.md) |
| slot\_id | mediumint | [Slot Identifier](https://eqemu.gitbook.io/server/categories/inventory/inventory-slots) |
| item\_id | int | [Item Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/bots/items.md) |
| inst\_charges | smallint | Charges |
| inst\_color | int | Color |
| inst\_no\_drop | tinyint | No Drop: 0 = False, 1=  True |
| inst\_custom\_data | text | Custom Data |
| ornament\_icon | int | Ornamentation Icon |
| ornament\_id\_file | int | Ornamentation Item Texture |
| ornament\_hero\_model | int | Ornamentation Hero's Forge Model |
| augment\_1 | mediumint | Augment Slot 1 |
| augment\_2 | mediumint | Augment Slot 2 |
| augment\_3 | mediumint | Augment Slot 3 |
| augment\_4 | mediumint | Augment Slot 4 |
| augment\_5 | mediumint | Augment Slot 5 |
| augment\_6 | mediumint | Augment Slot 6 |

