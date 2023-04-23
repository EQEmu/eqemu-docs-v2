# vw_guild_members

!!! info
	This page was last generated 2023.04.23

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| mob_type | varchar | Mob Type: B = Bot, C = Client |
| char_id | int | [Character Identifier](../../schema/characters/character_data.md) |
| guild_id | mediumint | [Guild Identifier](../../schema/guilds/guilds.md) |
| rank | tinyint | Ranl |
| tribute_enable | tinyint | Tribute Enable: 0 = False, 1 = True |
| total_tribute | int | Total Tribute |
| last_tribute | int | Last Tribute |
| banker | tinyint | Banker: 0 = False, 1 = True |
| public_note | mediumtext | Public Note |
| alt | tinyint | Alt: 0 = False, 1 = True |

