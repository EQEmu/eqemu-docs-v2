# qs_player_npc_kill_record_entries

!!! info
	This page was last generated 2022.06.19

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcXNfcGxheWVyX25wY19raWxsX3JlY29yZF9lbnRyaWVzIHtcbiAgICAgICAgaW50IGNoYXJfaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgfVxuICAgIHFzX3BsYXllcl9ucGNfa2lsbF9yZWNvcmRfZW50cmllcyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcXNfcGxheWVyX25wY19raWxsX3JlY29yZF9lbnRyaWVzIHtcbiAgICAgICAgaW50IGNoYXJfaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgfVxuICAgIHFzX3BsYXllcl9ucGNfa2lsbF9yZWNvcmRfZW50cmllcyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcXNfcGxheWVyX25wY19raWxsX3JlY29yZF9lbnRyaWVzIHtcbiAgICAgICAgaW50IGNoYXJfaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgfVxuICAgIHFzX3BsYXllcl9ucGNfa2lsbF9yZWNvcmRfZW50cmllcyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | char_id | [character_data](../../schema/characters/character_data.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| event_id | int | Unique Event Identifier |
| char_id | int | [Character Identifier](../../schema/characters/character_data.md) |

