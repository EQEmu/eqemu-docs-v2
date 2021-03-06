# mail

!!! info
	This page was last generated 2022.06.19

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWFpbCB7XG4gICAgICAgIGludHVuc2lnbmVkIGNoYXJpZFxuICAgICAgICB2YXJjaGFyIGZyb21cbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBtYWlsIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IE9uZS10by1PbmVcbiAgICBtYWlsIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWFpbCB7XG4gICAgICAgIGludHVuc2lnbmVkIGNoYXJpZFxuICAgICAgICB2YXJjaGFyIGZyb21cbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBtYWlsIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IE9uZS10by1PbmVcbiAgICBtYWlsIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWFpbCB7XG4gICAgICAgIGludHVuc2lnbmVkIGNoYXJpZFxuICAgICAgICB2YXJjaGFyIGZyb21cbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBtYWlsIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IE9uZS10by1PbmVcbiAgICBtYWlsIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | charid | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | from | [character_data](../../schema/characters/character_data.md) | name |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| msgid | int | Unique Message Identifier |
| charid | int | [Character Identifier](character_data.md) |
| timestamp | int | UNIX Timestamp |
| from | varchar | [From Character Name](character_data.md) |
| subject | varchar | Subject |
| body | text | Body |
| to | text | [To Character Name](character_data.md) |
| status | tinyint | [Status](../../../../server/player/status-levels) |

