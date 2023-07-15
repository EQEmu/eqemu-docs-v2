# spell_buckets

!!! info
	This page was last generated 2023.07.15

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3BlbGxfYnVja2V0cyB7XG4gICAgICAgIGJpZ2ludHVuc2lnbmVkIHNwZWxsaWRcbiAgICAgICAgdmFyY2hhciBrZXlcbiAgICB9XG4gICAgc3BlbGxzX25ldyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bTJcbiAgICAgICAgaW50IHR5cGVkZXNjbnVtXG4gICAgICAgIHZhcmNoYXIgdGVsZXBvcnRfem9uZVxuICAgIH1cbiAgICBkYXRhX2J1Y2tldHMge1xuICAgICAgICB2YXJjaGFyIGtleVxuICAgIH1cbiAgICBzcGVsbF9idWNrZXRzIHx8LS1veyBzcGVsbHNfbmV3IDogXCJPbmUtdG8tT25lXCJcbiAgICBzcGVsbF9idWNrZXRzIHx8LS1veyBkYXRhX2J1Y2tldHMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3BlbGxfYnVja2V0cyB7XG4gICAgICAgIGJpZ2ludHVuc2lnbmVkIHNwZWxsaWRcbiAgICAgICAgdmFyY2hhciBrZXlcbiAgICB9XG4gICAgc3BlbGxzX25ldyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bTJcbiAgICAgICAgaW50IHR5cGVkZXNjbnVtXG4gICAgICAgIHZhcmNoYXIgdGVsZXBvcnRfem9uZVxuICAgIH1cbiAgICBkYXRhX2J1Y2tldHMge1xuICAgICAgICB2YXJjaGFyIGtleVxuICAgIH1cbiAgICBzcGVsbF9idWNrZXRzIHx8LS1veyBzcGVsbHNfbmV3IDogXCJPbmUtdG8tT25lXCJcbiAgICBzcGVsbF9idWNrZXRzIHx8LS1veyBkYXRhX2J1Y2tldHMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3BlbGxfYnVja2V0cyB7XG4gICAgICAgIGJpZ2ludHVuc2lnbmVkIHNwZWxsaWRcbiAgICAgICAgdmFyY2hhciBrZXlcbiAgICB9XG4gICAgc3BlbGxzX25ldyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bTJcbiAgICAgICAgaW50IHR5cGVkZXNjbnVtXG4gICAgICAgIHZhcmNoYXIgdGVsZXBvcnRfem9uZVxuICAgIH1cbiAgICBkYXRhX2J1Y2tldHMge1xuICAgICAgICB2YXJjaGFyIGtleVxuICAgIH1cbiAgICBzcGVsbF9idWNrZXRzIHx8LS1veyBzcGVsbHNfbmV3IDogXCJPbmUtdG8tT25lXCJcbiAgICBzcGVsbF9idWNrZXRzIHx8LS1veyBkYXRhX2J1Y2tldHMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | spellid | [spells_new](../../schema/spells/spells_new.md) | id |
| One-to-One | key | [data_buckets](../../schema/data-storage/data_buckets.md) | key |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| spellid | bigint | [Unique Spell Identifier](spells_new.md) |
| key | varchar | [Data Bucket Name](../../schema/data-storage/data_buckets.md) |
| value | text | Data Bucket Value |

