# lootdrop\_entries

| Column | Data Type | Description |
| :--- | :--- | :--- |
| lootdrop\_id | int | [Lootdrop Identifier](lootdrop.md) |
| item\_id | int | [Item Identifier](../../../schema/categories/loot/items.md) |
| item\_charges | smallint | Item Charges |
| equip\_item | tinyint | Equip Item: 0 = False, 1 = True |
| chance | float | Chance: 0 = Never, 100 = Always |
| disabled\_chance | float | Disabled Chance: 0 = Never, 100 = Always |
| minlevel | tinyint | Minimum Level |
| maxlevel | tinyint | Maximum Level |
| multiplier | tinyint | Multiplier |

