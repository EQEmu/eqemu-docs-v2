# ldon\_trap\_templates

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique LDoN Trap Template Identifier |
| type | tinyint | [Trap Type](https://eqemu.gitbook.io/server/categories/zones/trap-types) |
| spell\_id | smallint | [Spell Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/traps/spells_new.md) |
| skill | smallint | [Skill](https://eqemu.gitbook.io/server/categories/player/skills) |
| locked | tinyint | Locked: 0 = False, 1 = True |

