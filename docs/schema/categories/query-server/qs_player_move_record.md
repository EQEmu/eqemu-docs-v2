# qs\_player\_move\_record

| Column | Data Type | Description |
| :--- | :--- | :--- |
| move\_id | int | Unique Move Identifier |
| time | timestamp | Time Timestamp |
| char\_id | int | [Character Identifier](https://github.com/EQEmu/docs-db-schema/tree/774e95edd473c84dafd6fe13b9b699f6b84a7ce8/docs/schema/categories/query_server/character_data.md) |
| from\_slot | mediumint | [From Slot Identifier](https://eqemu.gitbook.io/server/categories/inventory/inventory-slots) |
| to\_slot | mediumint | [To Slot Identifier](https://eqemu.gitbook.io/server/categories/inventory/inventory-slots) |
| stack\_size | mediumint | Stack Size |
| char\_items | mediumint | [Character Item Identifier](https://github.com/EQEmu/docs-db-schema/tree/774e95edd473c84dafd6fe13b9b699f6b84a7ce8/docs/schema/categories/query_server/items.md) |
| postaction | tinyint | Post Action |

