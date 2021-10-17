# ground\_spawns

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Ground Spawn Identifier |
| zoneid | int | [Zone Identifier](https://eqemu.gitbook.io/server/categories/zones/zone-list) |
| version | smallint | Version: -1 For All |
| max\_x | float | Maximum X Coordinate |
| max\_y | float | Maximum Y Coordinate |
| max\_z | float | Maximum Z Coordinate |
| min\_x | float | Minimum X Coordinate |
| min\_y | float | Minimum Y Coordinate |
| heading | float | Heading Coordinate |
| name | varchar | Name |
| item | int | [Item Identifier](https://github.com/EQEmu/docs-db-schema/tree/774e95edd473c84dafd6fe13b9b699f6b84a7ce8/docs/schema/categories/ground_spawns/items.md) |
| max\_allowed | int | Max Allowed |
| comment | varchar | Comment |
| respawn\_timer | int | Respawn Timer in Seconds |

