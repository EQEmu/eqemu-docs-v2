# character_buffs

!!! info
	This page was last generated 2023.06.12

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcmFjdGVyX2J1ZmZzIHtcbiAgICAgICAgaW50dW5zaWduZWQgY2hhcmFjdGVyX2lkXG4gICAgICAgIHNtYWxsaW50dW5zaWduZWQgc3BlbGxfaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgdmFyY2hhciBuYW5lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgIH1cbiAgICBzcGVsbHNfbmV3IHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBkZXNjbnVtXG4gICAgICAgIGludCBlZmZlY3RkZXNjbnVtXG4gICAgICAgIGludCBlZmZlY3RkZXNjbnVtMlxuICAgICAgICBpbnQgdHlwZWRlc2NudW1cbiAgICAgICAgdmFyY2hhciB0ZWxlcG9ydF96b25lXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9idWZmcyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBcIk9uZS10by1PbmVcIlxuICAgIGNoYXJhY3Rlcl9idWZmcyB8fC0tb3sgc3BlbGxzX25ldyA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcmFjdGVyX2J1ZmZzIHtcbiAgICAgICAgaW50dW5zaWduZWQgY2hhcmFjdGVyX2lkXG4gICAgICAgIHNtYWxsaW50dW5zaWduZWQgc3BlbGxfaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgdmFyY2hhciBuYW5lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgIH1cbiAgICBzcGVsbHNfbmV3IHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBkZXNjbnVtXG4gICAgICAgIGludCBlZmZlY3RkZXNjbnVtXG4gICAgICAgIGludCBlZmZlY3RkZXNjbnVtMlxuICAgICAgICBpbnQgdHlwZWRlc2NudW1cbiAgICAgICAgdmFyY2hhciB0ZWxlcG9ydF96b25lXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9idWZmcyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBcIk9uZS10by1PbmVcIlxuICAgIGNoYXJhY3Rlcl9idWZmcyB8fC0tb3sgc3BlbGxzX25ldyA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcmFjdGVyX2J1ZmZzIHtcbiAgICAgICAgaW50dW5zaWduZWQgY2hhcmFjdGVyX2lkXG4gICAgICAgIHNtYWxsaW50dW5zaWduZWQgc3BlbGxfaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgdmFyY2hhciBuYW5lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgIH1cbiAgICBzcGVsbHNfbmV3IHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBkZXNjbnVtXG4gICAgICAgIGludCBlZmZlY3RkZXNjbnVtXG4gICAgICAgIGludCBlZmZlY3RkZXNjbnVtMlxuICAgICAgICBpbnQgdHlwZWRlc2NudW1cbiAgICAgICAgdmFyY2hhciB0ZWxlcG9ydF96b25lXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9idWZmcyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBcIk9uZS10by1PbmVcIlxuICAgIGNoYXJhY3Rlcl9idWZmcyB8fC0tb3sgc3BlbGxzX25ldyA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | character_id | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | spell_id | [spells_new](../../schema/spells/spells_new.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| character_id | int | [Character Identifier](../../schema/characters/character_data.md) |
| slot_id | tinyint | Buff Slot |
| spell_id | smallint | [Buff Spell Identifier](../../schema/spells/spells_new.md) |
| caster_level | tinyint | Caster Level |
| caster_name | varchar | Caster Name |
| ticsremaining | int | Tics Remaining |
| counters | int | Counters |
| numhits | int | Number of Hits |
| melee_rune | int | Melee Rune |
| magic_rune | int | Magic Rune |
| persistent | tinyint | Persistent: 0 = False, 1 = True |
| dot_rune | int | Damage Over Time Rune |
| caston_x | int | X Coordinate |
| caston_y | int | Y Coordinate |
| caston_z | int | Z Coordinate |
| ExtraDIChance | int | Extra DI Chance |
| instrument_mod | int | Instrument Modifier |

