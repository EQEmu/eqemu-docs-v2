# titles

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Title Identifier |
| skill\_id | tinyint | [Skill Identifier](https://eqemu.gitbook.io/server/categories/player/skills) |
| min\_skill\_value | mediumint | Minimum Skill Value |
| max\_skill\_value | mediumint | Maximum Skill Value |
| min\_aa\_points | mediumint | Minimum AA Points |
| max\_aa\_points | mediumint | Maximum AA Points |
| class | tinyint | [Class](https://eqemu.gitbook.io/server/categories/player/class-list) |
| gender | tinyint | [Gender](https://eqemu.gitbook.io/server/categories/npc/genders) |
| char\_id | int | [Unique Character Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/titles/character_data.md) |
| status | int | [Required Status](https://eqemu.gitbook.io/server/categories/player/status-levels) |
| item\_id | int | [Item Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/titles/items.md) |
| prefix | varchar | Prefix |
| suffix | varchar | Suffix |
| title\_set | int | Title Set Identifier |

