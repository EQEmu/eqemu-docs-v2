# merc_name_types

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19uYW1lX3R5cGVzIHtcbiAgICAgICAgdmFyY2hhciBuYW1lX3R5cGVfaWRcbiAgICB9XG4gICAgbWVyY190ZW1wbGF0ZXMge1xuICAgICAgICB2YXJjaGFyIGRic3RyaW5nXG4gICAgICAgIHZhcmNoYXIgbWVyY190ZW1wbGF0ZV9pZFxuICAgICAgICB2YXJjaGFyIG5hbWVfdHlwZV9pZFxuICAgICAgICB2YXJjaGFyIG1lcmNfbnBjX3R5cGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX3N1YnR5cGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX3R5cGVfaWRcbiAgICB9XG4gICAgbWVyY19uYW1lX3R5cGVzIHx8LS1veyBtZXJjX3RlbXBsYXRlcyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19uYW1lX3R5cGVzIHtcbiAgICAgICAgdmFyY2hhciBuYW1lX3R5cGVfaWRcbiAgICB9XG4gICAgbWVyY190ZW1wbGF0ZXMge1xuICAgICAgICB2YXJjaGFyIGRic3RyaW5nXG4gICAgICAgIHZhcmNoYXIgbWVyY190ZW1wbGF0ZV9pZFxuICAgICAgICB2YXJjaGFyIG5hbWVfdHlwZV9pZFxuICAgICAgICB2YXJjaGFyIG1lcmNfbnBjX3R5cGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX3N1YnR5cGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX3R5cGVfaWRcbiAgICB9XG4gICAgbWVyY19uYW1lX3R5cGVzIHx8LS1veyBtZXJjX3RlbXBsYXRlcyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19uYW1lX3R5cGVzIHtcbiAgICAgICAgdmFyY2hhciBuYW1lX3R5cGVfaWRcbiAgICB9XG4gICAgbWVyY190ZW1wbGF0ZXMge1xuICAgICAgICB2YXJjaGFyIGRic3RyaW5nXG4gICAgICAgIHZhcmNoYXIgbWVyY190ZW1wbGF0ZV9pZFxuICAgICAgICB2YXJjaGFyIG5hbWVfdHlwZV9pZFxuICAgICAgICB2YXJjaGFyIG1lcmNfbnBjX3R5cGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX3N1YnR5cGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX3R5cGVfaWRcbiAgICB9XG4gICAgbWVyY19uYW1lX3R5cGVzIHx8LS1veyBtZXJjX3RlbXBsYXRlcyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | name_type_id | [merc_templates](../../schema/mercenaries/merc_templates.md) | name_type_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| name_type_id | int | Mercenary Name Type Identifier |
| class_id | int | [Class Identifier](../../../../server/player/class-list) |
| prefix | varchar | Prefix |
| suffix | varchar | Suffix |

