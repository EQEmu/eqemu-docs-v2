# qs\_player\_move\_record

| Column | Data Type | Description |
| :--- | :--- | :--- |
| move\_id | int | Unique Move Identifier |
| time | timestamp | Time Timestamp |
| char\_id | int | [Character Identifier](../../../schema/categories/query_server/character_data.md) |
| from\_slot | mediumint | [From Slot Identifier](../../../../categories/inventory/inventory-slots) |
| to\_slot | mediumint | [To Slot Identifier](../../../../categories/inventory/inventory-slots) |
| stack\_size | mediumint | Stack Size |
| char\_items | mediumint | [Character Item Identifier](../../../schema/categories/query_server/items.md) |
| postaction | tinyint | Post Action |

