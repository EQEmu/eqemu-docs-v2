# bot\_data

| Column | Data Type | Description |
| :--- | :--- | :--- |
| bot\_id | int | Unique Bot Identifier |
| owner\_id | int | [Owner Character Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/bots/character_data.md) |
| spells\_id | int | [Bot Spell List Identifier](https://eqemu.gitbook.io/server/categories/spells/bot-spell-list-ids) |
| name | varchar | Name |
| last\_name | varchar | Last Name |
| title | varchar | Title |
| suffix | varchar | Suffix |
| zone\_id | smallint | [Zone Identifier](https://eqemu.gitbook.io/server/categories/zones/zone-list) |
| gender | tinyint | [Gender](https://eqemu.gitbook.io/server/categories/npc/genders) |
| race | smallint | [Race](https://eqemu.gitbook.io/server/categories/npc/race-list) |
| class | tinyint | [Class](https://eqemu.gitbook.io/server/categories/player/class-list) |
| level | tinyint | Level |
| deity | int | [Deity](https://eqemu.gitbook.io/server/categories/player/deity-list) |
| creation\_day | int | UNIX Timestamp of creation date |
| last\_spawn | int | UNIX Timestamp of last spawn\_conditions |
| time\_spawned | int | Time spawned |
| size | float | Size |
| face | int | Face |
| hair\_color | int | Hair Color |
| hair\_style | int | Hair Style |
| beard | int | Beard |
| beard\_color | int | Beard Color |
| eye\_color\_1 | int | Eye Color 1 |
| eye\_color\_2 | int | Eye Color 2 |
| drakkin\_heritage | int | Drakkin Heritage |
| drakkin\_tattoo | int | Drakkin Tattoo |
| drakkin\_details | int | Drakkin Details |
| ac | smallint | Armor Class |
| atk | mediumint | Attack |
| hp | int | Health |
| mana | int | Mana |
| str | mediumint | Strength |
| sta | mediumint | Stamina |
| cha | mediumint | Charisma |
| dex | mediumint | Dexterity |
| int | mediumint | Intelligence |
| agi | mediumint | Agility |
| wis | mediumint | Wisdom |
| fire | smallint | Fire Resistance |
| cold | smallint | Cold Resistance |
| magic | smallint | Magic Resistance |
| poison | smallint | Poison Resistance |
| disease | smallint | Disease Resistance |
| corruption | smallint | Corruption Resistance |
| show\_helm | int | Show Helm: 0 = False, 1= True |
| follow\_distance | int | Follow Distance |
| stop\_melee\_level | tinyint | Stop Melee Level |

