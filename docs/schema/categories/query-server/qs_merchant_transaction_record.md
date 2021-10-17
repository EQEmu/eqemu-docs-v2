# qs\_merchant\_transaction\_record

| Column | Data Type | Description |
| :--- | :--- | :--- |
| transaction\_id | int | Unique Transaction Identifier |
| time | timestamp | Timestamp |
| zone\_id | int | [Zone Identifier](https://eqemu.gitbook.io/server/categories/zones/zone-list) |
| merchant\_id | int | [Merchant Identifier](https://github.com/EQEmu/docs-db-schema/tree/774e95edd473c84dafd6fe13b9b699f6b84a7ce8/docs/schema/categories/query_server/merchantlist.md) |
| merchant\_pp | int | Merchant Platinum |
| merchant\_gp | int | Merchant Gold |
| merchant\_sp | int | Merchant Silver |
| merchant\_cp | int | Merchant Copper |
| merchant\_items | mediumint | Merchant Items |
| char\_id | int | [Unique Character Identifier](https://github.com/EQEmu/docs-db-schema/tree/774e95edd473c84dafd6fe13b9b699f6b84a7ce8/docs/schema/categories/query_server/character_data.md) |
| char\_pp | int | Character Platinum |
| char\_gp | int | Character Gold |
| char\_sp | int | Character Silver |
| char\_cp | int | Character Copper |
| char\_items | mediumint | Character Items |

