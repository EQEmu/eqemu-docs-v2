# merc_buffs

!!! info
	This page was last generated 2022.06.19

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19idWZmcyB7XG4gICAgICAgIGludHVuc2lnbmVkIE1lcmNJZFxuICAgICAgICBpbnR1bnNpZ25lZCBTcGVsbElkXG4gICAgfVxuICAgIG1lcmMge1xuICAgICAgICB2YXJjaGFyIE1lcmNJRFxuICAgIH1cbiAgICBzcGVsbHNfbmV3IHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBkZXNjbnVtXG4gICAgICAgIGludCBlZmZlY3RkZXNjbnVtXG4gICAgICAgIGludCBlZmZlY3RkZXNjbnVtMlxuICAgICAgICBpbnQgdHlwZWRlc2NudW1cbiAgICAgICAgdmFyY2hhciB0ZWxlcG9ydF96b25lXG4gICAgfVxuICAgIG1lcmNfYnVmZnMgfHwtLW97IG1lcmMgOiBPbmUtdG8tT25lXG4gICAgbWVyY19idWZmcyB8fC0tb3sgc3BlbGxzX25ldyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19idWZmcyB7XG4gICAgICAgIGludHVuc2lnbmVkIE1lcmNJZFxuICAgICAgICBpbnR1bnNpZ25lZCBTcGVsbElkXG4gICAgfVxuICAgIG1lcmMge1xuICAgICAgICB2YXJjaGFyIE1lcmNJRFxuICAgIH1cbiAgICBzcGVsbHNfbmV3IHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBkZXNjbnVtXG4gICAgICAgIGludCBlZmZlY3RkZXNjbnVtXG4gICAgICAgIGludCBlZmZlY3RkZXNjbnVtMlxuICAgICAgICBpbnQgdHlwZWRlc2NudW1cbiAgICAgICAgdmFyY2hhciB0ZWxlcG9ydF96b25lXG4gICAgfVxuICAgIG1lcmNfYnVmZnMgfHwtLW97IG1lcmMgOiBPbmUtdG8tT25lXG4gICAgbWVyY19idWZmcyB8fC0tb3sgc3BlbGxzX25ldyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19idWZmcyB7XG4gICAgICAgIGludHVuc2lnbmVkIE1lcmNJZFxuICAgICAgICBpbnR1bnNpZ25lZCBTcGVsbElkXG4gICAgfVxuICAgIG1lcmMge1xuICAgICAgICB2YXJjaGFyIE1lcmNJRFxuICAgIH1cbiAgICBzcGVsbHNfbmV3IHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBkZXNjbnVtXG4gICAgICAgIGludCBlZmZlY3RkZXNjbnVtXG4gICAgICAgIGludCBlZmZlY3RkZXNjbnVtMlxuICAgICAgICBpbnQgdHlwZWRlc2NudW1cbiAgICAgICAgdmFyY2hhciB0ZWxlcG9ydF96b25lXG4gICAgfVxuICAgIG1lcmNfYnVmZnMgfHwtLW97IG1lcmMgOiBPbmUtdG8tT25lXG4gICAgbWVyY19idWZmcyB8fC0tb3sgc3BlbGxzX25ldyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | MercId | merc | MercID |
| One-to-One | SpellId | [spells_new](../../schema/spells/spells_new.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| MercBuffId | int | Unique Mercenary Buff Identifier |
| MercId | int | [Mercenary Identifier](mercs.md) |
| SpellId | int | [Spell Identifier](../../schema/spells/spells_new.md) |
| CasterLevel | int | Caster Level |
| DurationFormula | int | [Duration Formula](../../../../categories/spells/buff-duration-formulas) |
| TicsRemaining | int | Tics Remaining |
| PoisonCounters | int | Poison Counters |
| DiseaseCounters | int | Disease Counters |
| CurseCounters | int | Curse Counters |
| CorruptionCounters | int | Corruption Counters |
| HitCount | int | Hit Count |
| MeleeRune | int | Melee Rune |
| MagicRune | int | Magic Rune |
| dot_rune | int | Damage Over Time Rune |
| caston_x | int | Cast On X Coordinate |
| Persistent | tinyint | Persistent: 0 = False, 1 = True |
| caston_y | int | Cast On Y Coordinate |
| caston_z | int | Cast On Z Coordinate |
| ExtraDIChance | int | Extra DI Chance |

