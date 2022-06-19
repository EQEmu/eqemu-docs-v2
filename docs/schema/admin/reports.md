# reports

!!! info
	This page was last generated 2022.06.19

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcmVwb3J0cyB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgIH1cbiAgICBjaGFyYWN0ZXJfZGF0YSB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2lkXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICB9XG4gICAgcmVwb3J0cyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcmVwb3J0cyB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgIH1cbiAgICBjaGFyYWN0ZXJfZGF0YSB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2lkXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICB9XG4gICAgcmVwb3J0cyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcmVwb3J0cyB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgIH1cbiAgICBjaGFyYWN0ZXJfZGF0YSB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2lkXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICB9XG4gICAgcmVwb3J0cyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | name | [character_data](../../schema/characters/character_data.md) | name |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Report Identifier |
| name | varchar | Name |
| reported | varchar | Reported |
| reported_text | text | Reported Text |

