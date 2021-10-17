# bot\_spells\_entries

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Bot Spell Entry Identifier |
| npc\_spells\_id | int | [Bot Spell List Identifier](../../../../categories/spells/bot-spell-list-ids) |
| spellid | smallint | [Spell Identifier](../../../schema/categories/bots/spells_new.md) |
| type | int | [Spell Type](../../../../categories/spells/spell-types) |
| minlevel | tinyint | Minimum Level |
| maxlevel | tinyint | Maximum Level |
| manacost | smallint | Mana Cost |
| recast\_delay | int | Recast Delay |
| priority | smallint | Bot Spell Priority: Lower is better |
| resist\_adjust | int | Resist Adjustment |
| min\_hp | smallint | Minimum Health Percentage |
| max\_hp | smallint | Maximum Health Percentage |

