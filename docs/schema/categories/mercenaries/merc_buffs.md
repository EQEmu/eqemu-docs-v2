# merc\_buffs

| Column | Data Type | Description |
| :--- | :--- | :--- |
| MercBuffId | int | Unique Mercenary Buff Identifier |
| MercId | int | [Mercenary Identifier](mercs.md) |
| SpellId | int | [Spell Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/mercenaries/spells_new.md) |
| CasterLevel | int | Caster Level |
| DurationFormula | int | [Duration Formula](https://eqemu.gitbook.io/server/categories/spells/buff-duration-formulas) |
| TicsRemaining | int | Tics Remaining |
| PoisonCounters | int | Poison Counters |
| DiseaseCounters | int | Disease Counters |
| CurseCounters | int | Curse Counters |
| CorruptionCounters | int | Corruption Counters |
| HitCount | int | Hit Count |
| MeleeRune | int | Melee Rune |
| MagicRune | int | Magic Rune |
| dot\_rune | int | Damage Over Time Rune |
| caston\_x | int | Cast On X Coordinate |
| Persistent | tinyint | Persistent: 0 = False, 1 = True |
| caston\_y | int | Cast On Y Coordinate |
| caston\_z | int | Cast On Z Coordinate |
| ExtraDIChance | int | Extra DI Chance |

