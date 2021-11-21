# lootdrop_entries

| Column | Data Type | Description |
| :--- | :--- | :--- |
| minlevel | tinyint | Minimum Level |
| lootdrop_id | int | [Lootdrop Identifier](lootdrop.md) |
| item_id | int | [Item Identifier](../../../schema/categories/items/items.md) |
| item_charges | smallint | Item Charges |
| equip_item | tinyint | Equip Item: 0 = False, 1 = True |
| chance | float | Chance: 0 = Never, 100 = Always |
| disabled_chance | float | Disabled Chance: 0 = Never, 100 = Always |
| trivial_min_level | smallint |  |
| trivial_max_level | smallint |  |
| multiplier | tinyint | Multiplier |
| npc_min_level | smallint |  |
| npc_max_level | smallint |  |

