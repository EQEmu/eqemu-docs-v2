# titles

!!! info
	This page was last generated 2022.06.17

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Title Identifier |
| skill_id | tinyint | [Skill Identifier](../../../../server/player/skills) |
| min_skill_value | mediumint | Minimum Skill Value |
| max_skill_value | mediumint | Maximum Skill Value |
| min_aa_points | mediumint | Minimum AA Points |
| max_aa_points | mediumint | Maximum AA Points |
| class | tinyint | [Class](../../../../server/player/class-list) |
| gender | tinyint | [Gender](../../../../server/npc/genders) |
| char_id | int | [Unique Character Identifier](../../../schema/characters/character_data) |
| status | int | [Required Status](../../../../server/player/status-levels) |
| item_id | int | [Item Identifier](../../../schema/items/items) |
| prefix | varchar | Prefix |
| suffix | varchar | Suffix |
| title_set | int | Title Set Identifier |

