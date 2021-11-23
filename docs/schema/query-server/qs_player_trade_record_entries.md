# qs_player_trade_record_entries

!!! info
	This page was last generated 2021.11.23

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| event_id | int | Unique Event Identifier |
| from_id | int | [From Character Identifier](../../schema/characters/character_data.md) |
| from_slot | mediumint | [From Slot Identifier](../../../../server/inventory/inventory-slots) |
| to_id | int | [To Character Identifier](../../schema/characters/character_data.md) |
| to_slot | mediumint | [To Slot Identifier](../../../../server/inventory/inventory-slots) |
| item_id | int | [Item Identifier](../../schema/items/items.md) |
| charges | mediumint | Charges |
| aug_1 | int | Augment Slot 1 |
| aug_2 | int | Augment Slot 2 |
| aug_3 | int | Augment Slot 3 |
| aug_4 | int | Augment Slot 4 |
| aug_5 | int | Augment Slot 5 |

