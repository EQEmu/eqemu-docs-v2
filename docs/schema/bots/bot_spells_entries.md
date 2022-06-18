# bot_spells_entries

!!! info
	This page was last generated 2022.06.17

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Bot Spell Entry Identifier |
| npc_spells_id | int | [Bot Spell List Identifier](../../../../categories/spells/bot-spell-list-ids) |
| spellid | smallint | [Spell Identifier](../../../schema/categories/spells/spells_new.md) |
| type | int | [Spell Type](../../../../categories/spells/spell-types) |
| minlevel | tinyint | Minimum Level |
| maxlevel | tinyint | Maximum Level |
| manacost | smallint | Mana Cost |
| recast_delay | int | Recast Delay |
| priority | smallint | Bot Spell Priority: Lower is better |
| resist_adjust | int | Resist Adjustment |
| min_hp | smallint | Minimum Health Percentage |
| max_hp | smallint | Maximum Health Percentage |

