# character\_buffs

| Column | Data Type | Description |
| :--- | :--- | :--- |
| character\_id | int | [Character Identifier](character_data.md) |
| slot\_id | tinyint | Buff Slot |
| spell\_id | smallint | [Buff Spell Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/characters/spells_new.md) |
| caster\_level | tinyint | Caster Level |
| caster\_name | varchar | Caster Name |
| ticsremaining | int | Tics Remaining |
| counters | int | Counters |
| numhits | int | Number of Hits |
| melee\_rune | int | Melee Rune |
| magic\_rune | int | Magic Rune |
| persistent | tinyint | Persistent: 0 = False, 1 = True |
| dot\_rune | int | Damage Over Time Rune |
| caston\_x | int | X Coordinate |
| caston\_y | int | Y Coordinate |
| caston\_z | int | Z Coordinate |
| ExtraDIChance | int | Extra DI Chance |
| instrument\_mod | int | Instrument Modifier |

