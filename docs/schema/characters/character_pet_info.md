# character_pet_info

!!! info
	This page was last generated 2023.07.15

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcmFjdGVyX3BldF9pbmZvIHtcbiAgICAgICAgaW50IGNoYXJfaWRcbiAgICAgICAgaW50IHNwZWxsX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICB9XG4gICAgc3BlbGxzX25ldyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bTJcbiAgICAgICAgaW50IHR5cGVkZXNjbnVtXG4gICAgICAgIHZhcmNoYXIgdGVsZXBvcnRfem9uZVxuICAgIH1cbiAgICBjaGFyYWN0ZXJfcGV0X2luZm8gfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogXCJPbmUtdG8tT25lXCJcbiAgICBjaGFyYWN0ZXJfcGV0X2luZm8gfHwtLW97IHNwZWxsc19uZXcgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcmFjdGVyX3BldF9pbmZvIHtcbiAgICAgICAgaW50IGNoYXJfaWRcbiAgICAgICAgaW50IHNwZWxsX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICB9XG4gICAgc3BlbGxzX25ldyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bTJcbiAgICAgICAgaW50IHR5cGVkZXNjbnVtXG4gICAgICAgIHZhcmNoYXIgdGVsZXBvcnRfem9uZVxuICAgIH1cbiAgICBjaGFyYWN0ZXJfcGV0X2luZm8gfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogXCJPbmUtdG8tT25lXCJcbiAgICBjaGFyYWN0ZXJfcGV0X2luZm8gfHwtLW97IHNwZWxsc19uZXcgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcmFjdGVyX3BldF9pbmZvIHtcbiAgICAgICAgaW50IGNoYXJfaWRcbiAgICAgICAgaW50IHNwZWxsX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICB9XG4gICAgc3BlbGxzX25ldyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bTJcbiAgICAgICAgaW50IHR5cGVkZXNjbnVtXG4gICAgICAgIHZhcmNoYXIgdGVsZXBvcnRfem9uZVxuICAgIH1cbiAgICBjaGFyYWN0ZXJfcGV0X2luZm8gfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogXCJPbmUtdG8tT25lXCJcbiAgICBjaGFyYWN0ZXJfcGV0X2luZm8gfHwtLW97IHNwZWxsc19uZXcgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | char_id | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | spell_id | [spells_new](../../schema/spells/spells_new.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| char_id | int | [Character Identifier](character_data.md) |
| pet | int | Pet |
| petname | varchar | Pet Name |
| petpower | int | Pet Power |
| spell_id | int | [Spell Identifier](../../schema/spells/spells_new.md) |
| hp | int | Health |
| mana | int | Mana |
| size | float | Size |
| taunting | tinyint | Taunting: 0 = False, 1 = True |

