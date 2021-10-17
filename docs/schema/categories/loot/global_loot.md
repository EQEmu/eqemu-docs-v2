# global\_loot

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Global Loot Identifier |
| description | varchar | Description |
| loottable\_id | int | [Loottable Identifier](loottable.md) |
| enabled | tinyint | Enabled: 0 = False, 1 = True |
| min\_level | int | Minimum Level |
| max\_level | int | Maximum Level |
| rare | tinyint | Rare: 0 = False, 1 = True |
| raid | tinyint | Raid: 0 = False, 1 = True |
| race | mediumtext | [Race](https://eqemu.gitbook.io/server/categories/npc/race-list), multiple races supported if \| delimited |
| class | mediumtext | [Class](https://eqemu.gitbook.io/server/categories/player/class-list), multiple classes supported if \| delimited |
| bodytype | mediumtext | [Body Type](https://eqemu.gitbook.io/server/categories/npc/body-types), multiple body types supported if \| delimited |
| zone | mediumtext | [Zone Short Name](https://eqemu.gitbook.io/server/categories/zones/zone-list),, multiple zones supported if \| delimited |
| hot\_zone | tinyint |  |

