# bot\_buffs

| Column | Data Type | Description |
| :--- | :--- | :--- |
| buffs\_index | int | Unique Entry Identifier |
| bot\_id | int | [Bot Identifier](bot_data.md) |
| spell\_id | int | [Spell Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/bots/spells_new.md) |
| caster\_level | tinyint | Caster level |
| duration\_formula | int | Duration Formula |
| tics\_remaining | int | Tics Remaining |
| poison\_counters | int | Poison Counter |
| disease\_counters | int | Disease Counter |
| curse\_counters | int | Curse Counter |
| corruption\_counters | int | Corruption Counter |
| numhits | int | Number of Hits |
| melee\_rune | int | Melee Rune |
| magic\_rune | int | Magic Rune |
| dot\_rune | int | Damage Over Time Rune |
| persistent | tinyint | Persistent: 0 = False, 1 = True |
| caston\_x | int | X Coordinate |
| caston\_y | int | Y Coordinate |
| caston\_z | int | Z Coordinate |
| extra\_di\_chance | int | Extra DI Chance |
| instrument\_mod | int | Instrument Modifier |

