# reports

!!! info
	This page was last generated 2022.06.19

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcmVwb3J0cyB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgIH1cbiAgICBjaGFyYWN0ZXJfZGF0YSB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgICAgIHZhcmNoYXIgbmFuZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2lkXG4gICAgfVxuICAgIHJlcG9ydHMgfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcmVwb3J0cyB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgIH1cbiAgICBjaGFyYWN0ZXJfZGF0YSB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgICAgIHZhcmNoYXIgbmFuZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2lkXG4gICAgfVxuICAgIHJlcG9ydHMgfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcmVwb3J0cyB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgIH1cbiAgICBjaGFyYWN0ZXJfZGF0YSB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgICAgIHZhcmNoYXIgbmFuZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2lkXG4gICAgfVxuICAgIHJlcG9ydHMgfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


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

