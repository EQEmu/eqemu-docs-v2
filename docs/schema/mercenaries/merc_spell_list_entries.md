# merc_spell_list_entries

!!! info
	This page was last generated 2022.06.17

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| merc_spell_list_entry_id | int | Unique Mercenary Spell List Entry Identifier |
| merc_spell_list_id | int | [Mercenary Spell List Identifier](merc_spell_lists.md) |
| spell_id | int | [Spell Identifier](../../../schema/categories/spells/spells_new.md) |
| spell_type | int | [Spell Type](../../../../server/spells/spell-types) |
| stance_id | tinyint | [Stance Type Identifier](../../../../categories/bots/stance-types) |
| minlevel | tinyint | Minimum Level |
| maxlevel | tinyint | Maximum Level |
| slot | tinyint | Slot |
| procChance | tinyint | Proc Chance: 0 = Never, 100 = Always |

