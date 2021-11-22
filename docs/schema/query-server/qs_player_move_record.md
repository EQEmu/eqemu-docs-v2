# qs_player_move_record

| Column | Data Type | Description |
| :--- | :--- | :--- |
| move_id | int | Unique Move Identifier |
| time | timestamp | Time Timestamp |
| char_id | int | [Character Identifier](../../schema/characters/character_data.md) |
| from_slot | mediumint | [From Slot Identifier](../../../../server/inventory/inventory-slots) |
| to_slot | mediumint | [To Slot Identifier](../../../../server/inventory/inventory-slots) |
| stack_size | mediumint | Stack Size |
| char_items | mediumint | [Character Item Identifier](../../schema/items/items.md) |
| postaction | tinyint | Post Action |

