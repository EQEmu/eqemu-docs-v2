# character_pet_buffs

!!! info
	This page was last generated 2022.06.19

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcmFjdGVyX3BldF9idWZmcyB7XG4gICAgICAgIGludCBzcGVsbF9pZFxuICAgICAgICBpbnQgY2hhcl9pZFxuICAgIH1cbiAgICBjaGFyYWN0ZXJfZGF0YSB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgICAgIHZhcmNoYXIgbmFuZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2lkXG4gICAgfVxuICAgIHNwZWxsc19uZXcge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGRlc2NudW1cbiAgICAgICAgaW50IGVmZmVjdGRlc2NudW1cbiAgICAgICAgaW50IGVmZmVjdGRlc2NudW0yXG4gICAgICAgIGludCB0eXBlZGVzY251bVxuICAgICAgICB2YXJjaGFyIHRlbGVwb3J0X3pvbmVcbiAgICB9XG4gICAgY2hhcmFjdGVyX3BldF9idWZmcyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgY2hhcmFjdGVyX3BldF9idWZmcyB8fC0tb3sgc3BlbGxzX25ldyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcmFjdGVyX3BldF9idWZmcyB7XG4gICAgICAgIGludCBzcGVsbF9pZFxuICAgICAgICBpbnQgY2hhcl9pZFxuICAgIH1cbiAgICBjaGFyYWN0ZXJfZGF0YSB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgICAgIHZhcmNoYXIgbmFuZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2lkXG4gICAgfVxuICAgIHNwZWxsc19uZXcge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGRlc2NudW1cbiAgICAgICAgaW50IGVmZmVjdGRlc2NudW1cbiAgICAgICAgaW50IGVmZmVjdGRlc2NudW0yXG4gICAgICAgIGludCB0eXBlZGVzY251bVxuICAgICAgICB2YXJjaGFyIHRlbGVwb3J0X3pvbmVcbiAgICB9XG4gICAgY2hhcmFjdGVyX3BldF9idWZmcyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgY2hhcmFjdGVyX3BldF9idWZmcyB8fC0tb3sgc3BlbGxzX25ldyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcmFjdGVyX3BldF9idWZmcyB7XG4gICAgICAgIGludCBzcGVsbF9pZFxuICAgICAgICBpbnQgY2hhcl9pZFxuICAgIH1cbiAgICBjaGFyYWN0ZXJfZGF0YSB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgICAgIHZhcmNoYXIgbmFuZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2lkXG4gICAgfVxuICAgIHNwZWxsc19uZXcge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGRlc2NudW1cbiAgICAgICAgaW50IGVmZmVjdGRlc2NudW1cbiAgICAgICAgaW50IGVmZmVjdGRlc2NudW0yXG4gICAgICAgIGludCB0eXBlZGVzY251bVxuICAgICAgICB2YXJjaGFyIHRlbGVwb3J0X3pvbmVcbiAgICB9XG4gICAgY2hhcmFjdGVyX3BldF9idWZmcyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgY2hhcmFjdGVyX3BldF9idWZmcyB8fC0tb3sgc3BlbGxzX25ldyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


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
| slot | int | Slot |
| spell_id | int | [Spell Identifier](../../schema/spells/spells_new.md) |
| caster_level | tinyint | Caster Level |
| castername | varchar | Caster Name |
| ticsremaining | int | Tics Remaining |
| counters | int | Counters |
| numhits | int | Number of Hits |
| rune | int | Rune |
| instrument_mod | tinyint | Instrument Modifier |

