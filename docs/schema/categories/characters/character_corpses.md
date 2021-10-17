# character\_corpses

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Corpse Identifier |
| charid | int | [Character Identifier](character_data.md) |
| charname | varchar | Character Name |
| zone\_id | smallint | [Zone Identifier](https://eqemu.gitbook.io/server/categories/zones/zone-list) |
| instance\_id | smallint | Instance Identifier |
| x | float | X Coordinate |
| y | float | Y Coordinate |
| z | float | Z Coordinate |
| heading | float | Heading Coordinate |
| time\_of\_death | datetime | Time of Death |
| is\_rezzed | tinyint | Is Rezzed: 0 = False, 1 = True |
| is\_buried | tinyint | Is Buried: 0 = False, 1 = True |
| was\_at\_graveyard | tinyint | Was At Graveyard: 0 = False, 1 = True |
| is\_locked | tinyint | Is Locked: 0 = False, 1 = True |
| exp | int | Experience |
| size | int | Size |
| level | int | Level |
| race | int | [Race](https://eqemu.gitbook.io/server/categories/npc/race-list) |
| gender | int | [Gender](https://eqemu.gitbook.io/server/categories/npc/genders) |
| class | int | [Class](https://eqemu.gitbook.io/server/categories/player/class-list) |
| deity | int | [Deity](https://eqemu.gitbook.io/server/categories/player/deity-list) |
| texture | int | [Texture](https://eqemu.gitbook.io/server/categories/npc/textures) |
| helm\_texture | int | [Helm Texture](https://eqemu.gitbook.io/server/categories/npc/textures) |
| copper | int | Copper |
| silver | int | Silver |
| gold | int | Gold |
| platinum | int | Platinum |
| hair\_color | int | Hair Color |
| beard\_color | int | Beard Color |
| eye\_color\_1 | int | Eye Color 1 |
| eye\_color\_2 | int | Eye Color 2 |
| hair\_style | int | Hair Style |
| face | int | Face |
| beard | int | Beard |
| drakkin\_heritage | int | Drakkin Heritage |
| drakkin\_tattoo | int | Drakkin Tattoo |
| drakkin\_details | int | Drakkin Details |
| wc\_1 | int | Wear Change 1 |
| wc\_2 | int | Wear Change 2 |
| wc\_3 | int | Wear Change 3 |
| wc\_4 | int | Wear Change 4 |
| wc\_5 | int | Wear Change 5 |
| wc\_6 | int | Wear Change 6 |
| wc\_7 | int | Wear Change 7 |
| wc\_8 | int | Wear Change 8 |
| wc\_9 | int | Wear Change 9 |
| guild\_consent\_id | int |  |

