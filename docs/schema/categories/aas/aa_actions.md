# aa\_actions

| Column | Data Type | Description |
| :--- | :--- | :--- |
| aaid | mediumint | AA Identifier |
| rank | tinyint | Rank: Starts at 0 |
| reuse\_time | mediumint | Reuse timer in seconds |
| spell\_id | mediumint | [Spell Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/aas/spells_new.md) |
| target | tinyint | [AA Target Type](https://eqemu.gitbook.io/server/categories/aas/aa-target-types) |
| nonspell\_action | tinyint | [AA Nonspell Action](https://eqemu.gitbook.io/server/categories/aas/aa-nonspell-actions) |
| nonspell\_mana | mediumint | Mana that the nonspell action consumes. |
| nonspell\_duration | mediumint | Duration which may be used by the nonspell action. |
| redux\_aa | mediumint | The AA which reduces the reuse timer of the skill. |
| redux\_rate | tinyint | The multiplier of redux\_aa, as a percentage of total rate. |
| redux\_aa2 | mediumint | The second AA which reduces the reuse timer of the skill. |
| redux\_rate2 | tinyint | The multiplier of redux\_aa2, as a percentage of total rate. |

