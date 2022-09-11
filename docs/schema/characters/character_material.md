# character_material

!!! info
	This page was last generated 2022.09.10

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcmFjdGVyX21hdGVyaWFsIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBjaGFyYWN0ZXJfbWF0ZXJpYWwgfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcmFjdGVyX21hdGVyaWFsIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBjaGFyYWN0ZXJfbWF0ZXJpYWwgfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcmFjdGVyX21hdGVyaWFsIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBjaGFyYWN0ZXJfbWF0ZXJpYWwgfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | id | [character_data](../../schema/characters/character_data.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | [Character Identifier](character_data.md) |
| slot | tinyint | Slot |
| blue | tinyint | Blue: 0 = None, 255 = Max |
| green | tinyint | Green: 0 = None, 255 = Max |
| red | tinyint | Red: 0 = None, 255 = Max |
| use_tint | tinyint | Tint: 0 = None, 255 = MAx |
| color | int | Color |

