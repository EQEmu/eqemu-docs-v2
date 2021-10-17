# bot\_spells\_entries

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Bot Spell Entry Identifier |
| npc\_spells\_id | int | [Bot Spell List Identifier](https://eqemu.gitbook.io/server/categories/spells/bot-spell-list-ids) |
| spellid | smallint | [Spell Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/bots/spells_new.md) |
| type | int | [Spell Type](https://eqemu.gitbook.io/server/categories/spells/spell-types) |
| minlevel | tinyint | Minimum Level |
| maxlevel | tinyint | Maximum Level |
| manacost | smallint | Mana Cost |
| recast\_delay | int | Recast Delay |
| priority | smallint | Bot Spell Priority: Lower is better |
| resist\_adjust | int | Resist Adjustment |
| min\_hp | smallint | Minimum Health Percentage |
| max\_hp | smallint | Maximum Health Percentage |

