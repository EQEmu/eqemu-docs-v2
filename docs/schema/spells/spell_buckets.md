# spell_buckets

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3BlbGxfYnVja2V0cyB7XG4gICAgICAgIHZhcmNoYXIgcWdsb2JhbFxuICAgICAgICBiaWdpbnR1bnNpZ25lZCBzcGVsbGlkXG4gICAgfVxuICAgIHNwZWxsc19uZXcge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciB0ZWxlcG9ydF96b25lXG4gICAgICAgIGludCBkZXNjbnVtXG4gICAgICAgIGludCBlZmZlY3RkZXNjbnVtXG4gICAgICAgIGludCBlZmZlY3RkZXNjbnVtMlxuICAgICAgICBpbnQgdHlwZWRlc2NudW1cbiAgICB9XG4gICAgc3BlbGxfYnVja2V0cyB8fC0tb3sgc3BlbGxzX25ldyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3BlbGxfYnVja2V0cyB7XG4gICAgICAgIHZhcmNoYXIgcWdsb2JhbFxuICAgICAgICBiaWdpbnR1bnNpZ25lZCBzcGVsbGlkXG4gICAgfVxuICAgIHNwZWxsc19uZXcge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciB0ZWxlcG9ydF96b25lXG4gICAgICAgIGludCBkZXNjbnVtXG4gICAgICAgIGludCBlZmZlY3RkZXNjbnVtXG4gICAgICAgIGludCBlZmZlY3RkZXNjbnVtMlxuICAgICAgICBpbnQgdHlwZWRlc2NudW1cbiAgICB9XG4gICAgc3BlbGxfYnVja2V0cyB8fC0tb3sgc3BlbGxzX25ldyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3BlbGxfYnVja2V0cyB7XG4gICAgICAgIHZhcmNoYXIgcWdsb2JhbFxuICAgICAgICBiaWdpbnR1bnNpZ25lZCBzcGVsbGlkXG4gICAgfVxuICAgIHNwZWxsc19uZXcge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciB0ZWxlcG9ydF96b25lXG4gICAgICAgIGludCBkZXNjbnVtXG4gICAgICAgIGludCBlZmZlY3RkZXNjbnVtXG4gICAgICAgIGludCBlZmZlY3RkZXNjbnVtMlxuICAgICAgICBpbnQgdHlwZWRlc2NudW1cbiAgICB9XG4gICAgc3BlbGxfYnVja2V0cyB8fC0tb3sgc3BlbGxzX25ldyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | spellid | [spells_new](../../schema/spells/spells_new.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| spellid | bigint | [Unique Spell Identifier](spells_new.md) |
| key | varchar | [Data Bucket Name](../../schema/data-storage/data_buckets.md) |
| value | text | Data Bucket Value |

