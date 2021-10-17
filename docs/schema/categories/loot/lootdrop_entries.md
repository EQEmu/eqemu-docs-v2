# lootdrop\_entries

| Column | Data Type | Description |
| :--- | :--- | :--- |
| lootdrop\_id | int | [Lootdrop Identifier](lootdrop.md) |
| item\_id | int | [Item Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/loot/items.md) |
| item\_charges | smallint | Item Charges |
| equip\_item | tinyint | Equip Item: 0 = False, 1 = True |
| chance | float | Chance: 0 = Never, 100 = Always |
| disabled\_chance | float | Disabled Chance: 0 = Never, 100 = Always |
| minlevel | tinyint | Minimum Level |
| maxlevel | tinyint | Maximum Level |
| multiplier | tinyint | Multiplier |

