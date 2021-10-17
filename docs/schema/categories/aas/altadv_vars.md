# altadv\_vars

| Column | Data Type | Description |
| :--- | :--- | :--- |
| skill\_id | int | [AA Identifier](aa_ability.md) |
| name | varchar | Name |
| cost | int | Cost |
| max\_level | int | Maximum Level |
| hotkey\_sid | int | Hotkey SID |
| hotkey\_sid2 | int | Hotkey SID 2 |
| title\_sid | int | Title SID |
| desc\_sid | int | Description SID |
| type | tinyint | Type |
| spellid | int | [Spell Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/aas/spells_new.md) |
| prereq\_skill | int | Prerequisite Skill Level |
| prereq\_minpoints | int | Prerequisite Minimum Points |
| spell\_type | int | [Spell Type](https://eqemu.gitbook.io/server/categories/spells/spell-types) |
| spell\_refresh | int | Spell Refresh |
| classes | int | [Classes](https://eqemu.gitbook.io/server/categories/player/class-list) |
| berserker | int | Berseker: 0 = False, 1 = True |
| class\_type | int | Level |
| cost\_inc | tinyint | Cost Increase |
| aa\_expansion | smallint | [AA Expansion Identifier](https://eqemu.gitbook.io/server/categories/operation/expansion-list) |
| special\_category | int | Special Category |
| sof\_type | tinyint | [Secrets of Faydwer Type](https://eqemu.gitbook.io/server/categories/aas/aa-categories) |
| sof\_cost\_inc | tinyint | Secrets of Faydwer Cost Increase |
| sof\_max\_level | tinyint | Secrets of Faydwer Maximum Level |
| sof\_next\_skill | int | Secrets of Faydwer Next Skill |
| clientver | tinyint | [Client Version](https://eqemu.gitbook.io/server/categories/operation/expansion-list) |
| account\_time\_required | int | Account Time Required in Seconds |
| sof\_current\_level | tinyint | Secrets of Faydwer Current Level |
| sof\_next\_id | int | Secrets of Faydwer Next Identifier |
| level\_inc | tinyint | Secrets of Faydwer Level Increase |

