# eventlog

!!! info
	This page was last generated 2023.07.15

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZXZlbnRsb2cge1xuICAgICAgICB2YXJjaGFyIGFjY291bnRpZFxuICAgICAgICB2YXJjaGFyIGFjY291bnRuYW1lXG4gICAgICAgIHZhcmNoYXIgY2hhcm5hbWVcbiAgICB9XG4gICAgYWNjb3VudCB7XG4gICAgICAgIGludCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgICAgIHZhcmNoYXIgbmFuZVxuICAgIH1cbiAgICBldmVudGxvZyB8fC0tb3sgYWNjb3VudCA6IFwiT25lLXRvLU9uZVwiXG4gICAgZXZlbnRsb2cgfHwtLW97IGFjY291bnQgOiBcIk9uZS10by1PbmVcIlxuICAgIGV2ZW50bG9nIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZXZlbnRsb2cge1xuICAgICAgICB2YXJjaGFyIGFjY291bnRpZFxuICAgICAgICB2YXJjaGFyIGFjY291bnRuYW1lXG4gICAgICAgIHZhcmNoYXIgY2hhcm5hbWVcbiAgICB9XG4gICAgYWNjb3VudCB7XG4gICAgICAgIGludCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgICAgIHZhcmNoYXIgbmFuZVxuICAgIH1cbiAgICBldmVudGxvZyB8fC0tb3sgYWNjb3VudCA6IFwiT25lLXRvLU9uZVwiXG4gICAgZXZlbnRsb2cgfHwtLW97IGFjY291bnQgOiBcIk9uZS10by1PbmVcIlxuICAgIGV2ZW50bG9nIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZXZlbnRsb2cge1xuICAgICAgICB2YXJjaGFyIGFjY291bnRpZFxuICAgICAgICB2YXJjaGFyIGFjY291bnRuYW1lXG4gICAgICAgIHZhcmNoYXIgY2hhcm5hbWVcbiAgICB9XG4gICAgYWNjb3VudCB7XG4gICAgICAgIGludCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgICAgIHZhcmNoYXIgbmFuZVxuICAgIH1cbiAgICBldmVudGxvZyB8fC0tb3sgYWNjb3VudCA6IFwiT25lLXRvLU9uZVwiXG4gICAgZXZlbnRsb2cgfHwtLW97IGFjY291bnQgOiBcIk9uZS10by1PbmVcIlxuICAgIGV2ZW50bG9nIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | accountid | [account](../../schema/account/account.md) | id |
| One-to-One | accountname | [account](../../schema/account/account.md) | name |
| One-to-One | charname | [character_data](../../schema/characters/character_data.md) | name |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Event Identifier |
| accountname | varchar | [Account Name](../../schema/account/account.md) |
| accountid | int | [Account Identifier](../../schema/account/account.md) |
| status | int | [Status](../../../../server/player/status-levels) |
| charname | varchar | [Character Name](../../schema/characters/character_data.md) |
| target | varchar | Target |
| time | timestamp | TIme Timestamp |
| descriptiontype | varchar | Description Type |
| description | text | Description |
| event_nid | int | Event Identifier |

