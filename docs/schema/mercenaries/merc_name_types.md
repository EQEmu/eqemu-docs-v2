# merc_name_types

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19uYW1lX3R5cGVzIHtcbiAgICAgICAgdmFyY2hhciBuYW1lX3R5cGVfaWRcbiAgICB9XG4gICAgbWVyY190ZW1wbGF0ZXMge1xuICAgICAgICB2YXJjaGFyIGRic3RyaW5nXG4gICAgICAgIHZhcmNoYXIgbmFtZV90eXBlX2lkXG4gICAgICAgIHZhcmNoYXIgbWVyY19ucGNfdHlwZV9pZFxuICAgICAgICB2YXJjaGFyIG1lcmNfdGVtcGxhdGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX3R5cGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX3N1YnR5cGVfaWRcbiAgICB9XG4gICAgbWVyY19uYW1lX3R5cGVzIHx8LS1veyBtZXJjX3RlbXBsYXRlcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19uYW1lX3R5cGVzIHtcbiAgICAgICAgdmFyY2hhciBuYW1lX3R5cGVfaWRcbiAgICB9XG4gICAgbWVyY190ZW1wbGF0ZXMge1xuICAgICAgICB2YXJjaGFyIGRic3RyaW5nXG4gICAgICAgIHZhcmNoYXIgbmFtZV90eXBlX2lkXG4gICAgICAgIHZhcmNoYXIgbWVyY19ucGNfdHlwZV9pZFxuICAgICAgICB2YXJjaGFyIG1lcmNfdGVtcGxhdGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX3R5cGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX3N1YnR5cGVfaWRcbiAgICB9XG4gICAgbWVyY19uYW1lX3R5cGVzIHx8LS1veyBtZXJjX3RlbXBsYXRlcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19uYW1lX3R5cGVzIHtcbiAgICAgICAgdmFyY2hhciBuYW1lX3R5cGVfaWRcbiAgICB9XG4gICAgbWVyY190ZW1wbGF0ZXMge1xuICAgICAgICB2YXJjaGFyIGRic3RyaW5nXG4gICAgICAgIHZhcmNoYXIgbmFtZV90eXBlX2lkXG4gICAgICAgIHZhcmNoYXIgbWVyY19ucGNfdHlwZV9pZFxuICAgICAgICB2YXJjaGFyIG1lcmNfdGVtcGxhdGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX3R5cGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX3N1YnR5cGVfaWRcbiAgICB9XG4gICAgbWVyY19uYW1lX3R5cGVzIHx8LS1veyBtZXJjX3RlbXBsYXRlcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


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

