# bot_spell_settings

!!! info
	This page was last generated 2024.02.07

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Bot Spell Setting Identifier |
| bot_id | int | [Bot Identifier](bot_data.md) |
| spell_id | smallint | [Spell Identifier](../../schema/spells/spells_new.md) |
| priority | smallint | Priority |
| min_hp | smallint | Minimum Health Percentage |
| max_hp | smallint | Maximum Health Percentage |
| is_enabled | tinyint | Is Enabled: 0 = False, 1 = True |

