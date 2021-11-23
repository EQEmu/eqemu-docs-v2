# global_loot

## Schema
| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Global Loot Identifier |
| description | varchar | Description |
| loottable_id | int | [Loottable Identifier](loottable.md) |
| enabled | tinyint | Enabled: 0 = False, 1 = True |
| min_level | int | Minimum Level |
| max_level | int | Maximum Level |
| rare | tinyint | Rare: 0 = False, 1 = True |
| raid | tinyint | Raid: 0 = False, 1 = True |
| race | mediumtext | [Race](../../../../server/npc/race-list), multiple races supported if \ |
| class | mediumtext | [Class](../../../../server/player/class-list), multiple classes supported if \ |
| bodytype | mediumtext | [Body Type](../../../../server/npc/body-types), multiple body types supported if \ |
| zone | mediumtext | [Zone Short Name](../../../../server/zones/zone-list),, multiple zones supported if \ |
| hot_zone | tinyint |  |
| min_expansion | tinyint |  |
| max_expansion | tinyint |  |
| content_flags | varchar |  |
| content_flags_disabled | varchar |  |

