# qs_player_speech

!!! info
	This page was last generated 2023.07.15

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcXNfcGxheWVyX3NwZWVjaCB7XG4gICAgICAgIHZhcmNoYXIgZnJvbVxuICAgICAgICB2YXJjaGFyIHRvXG4gICAgICAgIGludCBndWlsZGRiaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgICAgIHZhcmNoYXIgbmFuZVxuICAgIH1cbiAgICBndWlsZHMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGxlYWRlclxuICAgICAgICB2YXJjaGFyIG1vdGRfc2V0dGVyXG4gICAgfVxuICAgIHFzX3BsYXllcl9zcGVlY2ggfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogXCJPbmUtdG8tT25lXCJcbiAgICBxc19wbGF5ZXJfc3BlZWNoIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IFwiT25lLXRvLU9uZVwiXG4gICAgcXNfcGxheWVyX3NwZWVjaCB8fC0tb3sgZ3VpbGRzIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcXNfcGxheWVyX3NwZWVjaCB7XG4gICAgICAgIHZhcmNoYXIgZnJvbVxuICAgICAgICB2YXJjaGFyIHRvXG4gICAgICAgIGludCBndWlsZGRiaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgICAgIHZhcmNoYXIgbmFuZVxuICAgIH1cbiAgICBndWlsZHMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGxlYWRlclxuICAgICAgICB2YXJjaGFyIG1vdGRfc2V0dGVyXG4gICAgfVxuICAgIHFzX3BsYXllcl9zcGVlY2ggfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogXCJPbmUtdG8tT25lXCJcbiAgICBxc19wbGF5ZXJfc3BlZWNoIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IFwiT25lLXRvLU9uZVwiXG4gICAgcXNfcGxheWVyX3NwZWVjaCB8fC0tb3sgZ3VpbGRzIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcXNfcGxheWVyX3NwZWVjaCB7XG4gICAgICAgIHZhcmNoYXIgZnJvbVxuICAgICAgICB2YXJjaGFyIHRvXG4gICAgICAgIGludCBndWlsZGRiaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgICAgIHZhcmNoYXIgbmFuZVxuICAgIH1cbiAgICBndWlsZHMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGxlYWRlclxuICAgICAgICB2YXJjaGFyIG1vdGRfc2V0dGVyXG4gICAgfVxuICAgIHFzX3BsYXllcl9zcGVlY2ggfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogXCJPbmUtdG8tT25lXCJcbiAgICBxc19wbGF5ZXJfc3BlZWNoIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IFwiT25lLXRvLU9uZVwiXG4gICAgcXNfcGxheWVyX3NwZWVjaCB8fC0tb3sgZ3VpbGRzIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | from | [character_data](../../schema/characters/character_data.md) | name |
| One-to-One | to | [character_data](../../schema/characters/character_data.md) | name |
| One-to-One | guilddbid | [guilds](../../schema/guilds/guilds.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Speech Identifier |
| from | varchar | [From Character Identifier](../../schema/characters/character_data.md) |
| to | varchar | [To Character Identifier](../../schema/characters/character_data.md) |
| message | varchar | Message |
| minstatus | smallint | [Minimum Status](../../../../server/player/status-levels) |
| guilddbid | int | [Guild Database Identifier](../../schema/guilds/guilds.md) |
| type | tinyint | Type |
| timerecorded | timestamp | Time Recorded Timestamp |

