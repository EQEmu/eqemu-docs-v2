# inventory_snapshots

!!! info
	This page was last generated 2021.11.23

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| time_index | int | Inventory Snapshot Time Identifier |
| charid | int | [Character Identifier](../../schema/characters/character_data.md) |
| slotid | mediumint | [Slot Identifier](../../../../server/inventory/inventory-slots) |
| itemid | int | [Item Identifier](../../schema/items/items.md) |
| charges | smallint | Charges |
| color | int | Color |
| augslot1 | mediumint | Augment Slot 1 |
| augslot2 | mediumint | Augment Slot 2 |
| augslot3 | mediumint | Augment Slot 3 |
| augslot4 | mediumint | Augment Slot 4 |
| augslot5 | mediumint | Augment Slot 5 |
| augslot6 | mediumint | Augment Slot 6 |
| instnodrop | tinyint | No Drop: 0 = False, 1 = True |
| custom_data | text | Custom Data |
| ornamenticon | int | Ornamentation Icon |
| ornamentidfile | int | Ornamentation Texture |
| ornament_hero_model | int | Ornamentation Hero's Forge Model |

