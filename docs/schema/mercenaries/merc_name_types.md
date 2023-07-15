# merc_name_types

!!! info
	This page was last generated 2023.07.15

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19uYW1lX3R5cGVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgbmFtZV90eXBlX2lkXG4gICAgfVxuICAgIG1lcmNfdGVtcGxhdGVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY190eXBlX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfc3VidHlwZV9pZFxuICAgICAgICB2YXJjaGFyIGRic3RyaW5nXG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfdGVtcGxhdGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY19ucGNfdHlwZV9pZFxuICAgICAgICB0aW55aW50IG5hbWVfdHlwZV9pZFxuICAgIH1cbiAgICBtZXJjX25hbWVfdHlwZXMgfHwtLW97IG1lcmNfdGVtcGxhdGVzIDogXCJIYXMtTWFueVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19uYW1lX3R5cGVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgbmFtZV90eXBlX2lkXG4gICAgfVxuICAgIG1lcmNfdGVtcGxhdGVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY190eXBlX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfc3VidHlwZV9pZFxuICAgICAgICB2YXJjaGFyIGRic3RyaW5nXG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfdGVtcGxhdGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY19ucGNfdHlwZV9pZFxuICAgICAgICB0aW55aW50IG5hbWVfdHlwZV9pZFxuICAgIH1cbiAgICBtZXJjX25hbWVfdHlwZXMgfHwtLW97IG1lcmNfdGVtcGxhdGVzIDogXCJIYXMtTWFueVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19uYW1lX3R5cGVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgbmFtZV90eXBlX2lkXG4gICAgfVxuICAgIG1lcmNfdGVtcGxhdGVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY190eXBlX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfc3VidHlwZV9pZFxuICAgICAgICB2YXJjaGFyIGRic3RyaW5nXG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfdGVtcGxhdGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY19ucGNfdHlwZV9pZFxuICAgICAgICB0aW55aW50IG5hbWVfdHlwZV9pZFxuICAgIH1cbiAgICBtZXJjX25hbWVfdHlwZXMgfHwtLW97IG1lcmNfdGVtcGxhdGVzIDogXCJIYXMtTWFueVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | name_type_id | [merc_templates](../../schema/mercenaries/merc_templates.md) | name_type_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| name_type_id | int | Mercenary Name Type Identifier |
| class_id | int | [Class Identifier](../../../../categories/player/class-list) |
| prefix | varchar | Prefix |
| suffix | varchar | Suffix |

