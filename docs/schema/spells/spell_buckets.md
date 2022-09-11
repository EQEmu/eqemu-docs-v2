# spell_buckets

!!! info
	This page was last generated 2022.09.10

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3BlbGxfYnVja2V0cyB7XG4gICAgICAgIHZhcmNoYXIga2V5XG4gICAgICAgIGJpZ2ludHVuc2lnbmVkIHNwZWxsaWRcbiAgICB9XG4gICAgc3BlbGxzX25ldyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bTJcbiAgICAgICAgaW50IHR5cGVkZXNjbnVtXG4gICAgICAgIHZhcmNoYXIgdGVsZXBvcnRfem9uZVxuICAgIH1cbiAgICBkYXRhX2J1Y2tldHMge1xuICAgICAgICB2YXJjaGFyIGtleVxuICAgIH1cbiAgICBzcGVsbF9idWNrZXRzIHx8LS1veyBzcGVsbHNfbmV3IDogT25lLXRvLU9uZVxuICAgIHNwZWxsX2J1Y2tldHMgfHwtLW97IGRhdGFfYnVja2V0cyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3BlbGxfYnVja2V0cyB7XG4gICAgICAgIHZhcmNoYXIga2V5XG4gICAgICAgIGJpZ2ludHVuc2lnbmVkIHNwZWxsaWRcbiAgICB9XG4gICAgc3BlbGxzX25ldyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bTJcbiAgICAgICAgaW50IHR5cGVkZXNjbnVtXG4gICAgICAgIHZhcmNoYXIgdGVsZXBvcnRfem9uZVxuICAgIH1cbiAgICBkYXRhX2J1Y2tldHMge1xuICAgICAgICB2YXJjaGFyIGtleVxuICAgIH1cbiAgICBzcGVsbF9idWNrZXRzIHx8LS1veyBzcGVsbHNfbmV3IDogT25lLXRvLU9uZVxuICAgIHNwZWxsX2J1Y2tldHMgfHwtLW97IGRhdGFfYnVja2V0cyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3BlbGxfYnVja2V0cyB7XG4gICAgICAgIHZhcmNoYXIga2V5XG4gICAgICAgIGJpZ2ludHVuc2lnbmVkIHNwZWxsaWRcbiAgICB9XG4gICAgc3BlbGxzX25ldyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bTJcbiAgICAgICAgaW50IHR5cGVkZXNjbnVtXG4gICAgICAgIHZhcmNoYXIgdGVsZXBvcnRfem9uZVxuICAgIH1cbiAgICBkYXRhX2J1Y2tldHMge1xuICAgICAgICB2YXJjaGFyIGtleVxuICAgIH1cbiAgICBzcGVsbF9idWNrZXRzIHx8LS1veyBzcGVsbHNfbmV3IDogT25lLXRvLU9uZVxuICAgIHNwZWxsX2J1Y2tldHMgfHwtLW97IGRhdGFfYnVja2V0cyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


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

