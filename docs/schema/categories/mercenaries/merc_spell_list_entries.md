# merc\_spell\_list\_entries

| Column | Data Type | Description |
| :--- | :--- | :--- |
| merc\_spell\_list\_entry\_id | int | Unique Mercenary Spell List Entry Identifier |
| merc\_spell\_list\_id | int | [Mercenary Spell List Identifier](merc_spell_lists.md) |
| spell\_id | int | [Spell Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/mercenaries/spells_new.md) |
| spell\_type | int | [Spell Type](https://eqemu.gitbook.io/server/categories/spells/spell-types) |
| stance\_id | tinyint | [Stance Type Identifier](https://eqemu.gitbook.io/server/categories/bots/stance-types) |
| minlevel | tinyint | Minimum Level |
| maxlevel | tinyint | Maximum Level |
| slot | tinyint | Slot |
| procChance | tinyint | Proc Chance: 0 = Never, 100 = Always |

