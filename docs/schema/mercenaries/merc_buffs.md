# merc_buffs

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19idWZmcyB7XG4gICAgICAgIHZhcmNoYXIgTWVyY0lkXG4gICAgICAgIHZhcmNoYXIgU3BlbGxJZFxuICAgIH1cbiAgICBtZXJjIHtcbiAgICAgICAgdmFyY2hhciBNZXJjSURcbiAgICB9XG4gICAgc3BlbGxzX25ldyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICB2YXJjaGFyIHRlbGVwb3J0X3pvbmVcbiAgICAgICAgaW50IGRlc2NudW1cbiAgICAgICAgaW50IGVmZmVjdGRlc2NudW1cbiAgICAgICAgaW50IGVmZmVjdGRlc2NudW0yXG4gICAgICAgIGludCB0eXBlZGVzY251bVxuICAgIH1cbiAgICBtZXJjX2J1ZmZzIHx8LS1veyBtZXJjIDogT25lLXRvLU9uZVxuICAgIG1lcmNfYnVmZnMgfHwtLW97IHNwZWxsc19uZXcgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19idWZmcyB7XG4gICAgICAgIHZhcmNoYXIgTWVyY0lkXG4gICAgICAgIHZhcmNoYXIgU3BlbGxJZFxuICAgIH1cbiAgICBtZXJjIHtcbiAgICAgICAgdmFyY2hhciBNZXJjSURcbiAgICB9XG4gICAgc3BlbGxzX25ldyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICB2YXJjaGFyIHRlbGVwb3J0X3pvbmVcbiAgICAgICAgaW50IGRlc2NudW1cbiAgICAgICAgaW50IGVmZmVjdGRlc2NudW1cbiAgICAgICAgaW50IGVmZmVjdGRlc2NudW0yXG4gICAgICAgIGludCB0eXBlZGVzY251bVxuICAgIH1cbiAgICBtZXJjX2J1ZmZzIHx8LS1veyBtZXJjIDogT25lLXRvLU9uZVxuICAgIG1lcmNfYnVmZnMgfHwtLW97IHNwZWxsc19uZXcgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19idWZmcyB7XG4gICAgICAgIHZhcmNoYXIgTWVyY0lkXG4gICAgICAgIHZhcmNoYXIgU3BlbGxJZFxuICAgIH1cbiAgICBtZXJjIHtcbiAgICAgICAgdmFyY2hhciBNZXJjSURcbiAgICB9XG4gICAgc3BlbGxzX25ldyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICB2YXJjaGFyIHRlbGVwb3J0X3pvbmVcbiAgICAgICAgaW50IGRlc2NudW1cbiAgICAgICAgaW50IGVmZmVjdGRlc2NudW1cbiAgICAgICAgaW50IGVmZmVjdGRlc2NudW0yXG4gICAgICAgIGludCB0eXBlZGVzY251bVxuICAgIH1cbiAgICBtZXJjX2J1ZmZzIHx8LS1veyBtZXJjIDogT25lLXRvLU9uZVxuICAgIG1lcmNfYnVmZnMgfHwtLW97IHNwZWxsc19uZXcgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


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
| SpellId | int | [Spell Identifier](../../../schema/spells/spells_new.md) |
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

