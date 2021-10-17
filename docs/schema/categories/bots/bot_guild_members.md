# bot\_guild\_members

| Column | Data Type | Description |
| :--- | :--- | :--- |
| bot\_id | int | [Bot Identifier](bot_data.md) |
| guild\_id | mediumint | [Guild Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/bots/guilds.md) |
| rank | tinyint | [Guild Rank](https://eqemu.gitbook.io/server/categories/player/guild-ranks) |
| tribute\_enable | tinyint | Tribute Enabled: 0 = False, 1= True |
| total\_tribute | int | Total Tribute |
| last\_tribute | int | Last Tribute |
| banker | tinyint | Banker: 0 = False, 1 = True |
| public\_note | text | Public Note |
| alt | tinyint | Alt: 0 = False, 1 = True |

