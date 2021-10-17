# guild\_members

| Column | Data Type | Description |
| :--- | :--- | :--- |
| char\_id | int | [Character Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/guilds/character_data.md) |
| guild\_id | mediumint | [Guild Identifier](guilds.md) |
| rank | tinyint | [Rank](https://eqemu.gitbook.io/server/categories/player/guild-ranks) |
| tribute\_enable | tinyint | Tribute Enable: 0 = False, 1 = True |
| total\_tribute | int | Total Tribute |
| last\_tribute | int | Last Tribute |
| banker | tinyint | Banked: 0 = False, 1 = True |
| public\_note | text | Public Note |
| alt | tinyint | Alt: 0 = False, 1 = True |

