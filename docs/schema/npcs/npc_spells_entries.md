# npc_spells_entries

!!! info
	This page was last generated 2022.06.17

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique NPC Spell Entry Identifier |
| npc_spells_id | int | [Unique NPC Spell Set Identifier](npc_spells.md) |
| spellid | smallint | [Spell Identifier](spells_new.md) |
| type | int | [Spell Type Bitmask](../../../../server/spells/spell-types) |
| minlevel | tinyint | Minimum Level |
| maxlevel | tinyint | Maximum Level |
| manacost | smallint | Mana Cost |
| recast_delay | int | Recast Delay |
| priority | smallint | Priority: 0 = Innate, 1 = Highest Priority, 5 = Lower Priority, 10 = Even Lower Priority |
| resist_adjust | int | Resist Adjustment |
| min_hp | smallint | Minimum Health Percentage |
| max_hp | smallint | Maximum Health Percentage |

