# forage

## Schema
| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Forage Identifier |
| zoneid | int | [Zone Identifier](../../../../server/zones/zone-list) |
| Itemid | int | [Item Identifier](../../schema/items/items.md) |
| level | smallint | Level |
| chance | smallint | Chance: 0 = Never, 100 = Always |
| min_expansion | tinyint |  |
| max_expansion | tinyint |  |
| content_flags | varchar |  |
| content_flags_disabled | varchar |  |

