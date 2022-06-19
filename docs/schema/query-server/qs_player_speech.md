# qs_player_speech

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcXNfcGxheWVyX3NwZWVjaCB7XG4gICAgICAgIHZhcmNoYXIgZnJvbVxuICAgICAgICB2YXJjaGFyIHRvXG4gICAgICAgIGludCBndWlsZGRiaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2lkXG4gICAgfVxuICAgIGd1aWxkcyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgbGVhZGVyXG4gICAgICAgIHZhcmNoYXIgbW90ZF9zZXR0ZXJcbiAgICB9XG4gICAgcXNfcGxheWVyX3NwZWVjaCB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgcXNfcGxheWVyX3NwZWVjaCB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgcXNfcGxheWVyX3NwZWVjaCB8fC0tb3sgZ3VpbGRzIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcXNfcGxheWVyX3NwZWVjaCB7XG4gICAgICAgIHZhcmNoYXIgZnJvbVxuICAgICAgICB2YXJjaGFyIHRvXG4gICAgICAgIGludCBndWlsZGRiaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2lkXG4gICAgfVxuICAgIGd1aWxkcyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgbGVhZGVyXG4gICAgICAgIHZhcmNoYXIgbW90ZF9zZXR0ZXJcbiAgICB9XG4gICAgcXNfcGxheWVyX3NwZWVjaCB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgcXNfcGxheWVyX3NwZWVjaCB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgcXNfcGxheWVyX3NwZWVjaCB8fC0tb3sgZ3VpbGRzIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcXNfcGxheWVyX3NwZWVjaCB7XG4gICAgICAgIHZhcmNoYXIgZnJvbVxuICAgICAgICB2YXJjaGFyIHRvXG4gICAgICAgIGludCBndWlsZGRiaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2lkXG4gICAgfVxuICAgIGd1aWxkcyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgbGVhZGVyXG4gICAgICAgIHZhcmNoYXIgbW90ZF9zZXR0ZXJcbiAgICB9XG4gICAgcXNfcGxheWVyX3NwZWVjaCB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgcXNfcGxheWVyX3NwZWVjaCB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgcXNfcGxheWVyX3NwZWVjaCB8fC0tb3sgZ3VpbGRzIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


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

