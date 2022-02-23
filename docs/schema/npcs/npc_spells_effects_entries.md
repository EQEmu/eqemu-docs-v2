# npc_spells_effects_entries

!!! info
	This page was last generated 2022.02.23

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Spell Effect Entry Identifier |
| npc_spells_effects_id | int | [NPC Spells Effects Identifier](npc_spells_effects.md) |
| spell_effect_id | smallint | [Spell Effect Identifier](../../../../server/spells/spell-effect-ids) |
| minlevel | tinyint | Minimum Level |
| maxlevel | tinyint | Maximum Level |
| se_base | int | Spell Effect Base |
| se_limit | int | Spell Effect Limit |
| se_max | int | Spell Effect Maximum |

