# npc\_spells\_entries

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique NPC Spell Entry Identifier |
| npc\_spells\_id | int | [Unique NPC Spell Set Identifier](npc_spells.md) |
| spellid | smallint | [Spell Identifier](../../../schema/categories/npcs/spells_new.md) |
| type | int | [Spell Type Bitmask](../../../../categories/spells/spell-types) |
| minlevel | tinyint | Minimum Level |
| maxlevel | tinyint | Maximum Level |
| manacost | smallint | Mana Cost |
| recast\_delay | int | Recast Delay |
| priority | smallint | Priority: 0 = Innate, 1 = Highest Priority, 5 = Lower Priority, 10 = Even Lower Priority |
| resist\_adjust | int | Resist Adjustment |
| min\_hp | smallint | Minimum Health Percentage |
| max\_hp | smallint | Maximum Health Percentage |

