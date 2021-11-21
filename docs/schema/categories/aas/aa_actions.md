# aa_actions

| Column | Data Type | Description |
| :--- | :--- | :--- |
| aaid | mediumint | AA Identifier |
| rank | tinyint | Rank: Starts at 0 |
| reuse_time | mediumint | Reuse timer in seconds |
| spell_id | mediumint | [Spell Identifier](../../../schema/categories/spells/spells_new.md) |
| target | tinyint | [AA Target Type](../../../../categories/aas/aa-target-types) |
| nonspell_action | tinyint | [AA Nonspell Action](../../../../categories/aas/aa-nonspell-actions) |
| nonspell_mana | mediumint | Mana that the nonspell action consumes. |
| nonspell_duration | mediumint | Duration which may be used by the nonspell action. |
| redux_aa | mediumint | The AA which reduces the reuse timer of the skill. |
| redux_rate | tinyint | The multiplier of redux_aa, as a percentage of total rate. |
| redux_aa2 | mediumint | The second AA which reduces the reuse timer of the skill. |
| redux_rate2 | tinyint | The multiplier of redux_aa2, as a percentage of total rate. |

