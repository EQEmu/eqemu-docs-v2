# merc_types

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY190eXBlcyB7XG4gICAgICAgIHZhcmNoYXIgZGJzdHJpbmdcbiAgICAgICAgdmFyY2hhciBtZXJjX3R5cGVfaWRcbiAgICB9XG4gICAgZGJfc3RyIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIG1lcmNfdGVtcGxhdGVzIHtcbiAgICAgICAgdmFyY2hhciBkYnN0cmluZ1xuICAgICAgICB2YXJjaGFyIG5hbWVfdHlwZV9pZFxuICAgICAgICB2YXJjaGFyIG1lcmNfbnBjX3R5cGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX3RlbXBsYXRlX2lkXG4gICAgICAgIHZhcmNoYXIgbWVyY190eXBlX2lkXG4gICAgICAgIHZhcmNoYXIgbWVyY19zdWJ0eXBlX2lkXG4gICAgfVxuICAgIG1lcmNfdHlwZXMgfHwtLW97IGRiX3N0ciA6IE9uZS10by1PbmVcbiAgICBtZXJjX3R5cGVzIHx8LS1veyBtZXJjX3RlbXBsYXRlcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY190eXBlcyB7XG4gICAgICAgIHZhcmNoYXIgZGJzdHJpbmdcbiAgICAgICAgdmFyY2hhciBtZXJjX3R5cGVfaWRcbiAgICB9XG4gICAgZGJfc3RyIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIG1lcmNfdGVtcGxhdGVzIHtcbiAgICAgICAgdmFyY2hhciBkYnN0cmluZ1xuICAgICAgICB2YXJjaGFyIG5hbWVfdHlwZV9pZFxuICAgICAgICB2YXJjaGFyIG1lcmNfbnBjX3R5cGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX3RlbXBsYXRlX2lkXG4gICAgICAgIHZhcmNoYXIgbWVyY190eXBlX2lkXG4gICAgICAgIHZhcmNoYXIgbWVyY19zdWJ0eXBlX2lkXG4gICAgfVxuICAgIG1lcmNfdHlwZXMgfHwtLW97IGRiX3N0ciA6IE9uZS10by1PbmVcbiAgICBtZXJjX3R5cGVzIHx8LS1veyBtZXJjX3RlbXBsYXRlcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY190eXBlcyB7XG4gICAgICAgIHZhcmNoYXIgZGJzdHJpbmdcbiAgICAgICAgdmFyY2hhciBtZXJjX3R5cGVfaWRcbiAgICB9XG4gICAgZGJfc3RyIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIG1lcmNfdGVtcGxhdGVzIHtcbiAgICAgICAgdmFyY2hhciBkYnN0cmluZ1xuICAgICAgICB2YXJjaGFyIG5hbWVfdHlwZV9pZFxuICAgICAgICB2YXJjaGFyIG1lcmNfbnBjX3R5cGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX3RlbXBsYXRlX2lkXG4gICAgICAgIHZhcmNoYXIgbWVyY190eXBlX2lkXG4gICAgICAgIHZhcmNoYXIgbWVyY19zdWJ0eXBlX2lkXG4gICAgfVxuICAgIG1lcmNfdHlwZXMgfHwtLW97IGRiX3N0ciA6IE9uZS10by1PbmVcbiAgICBtZXJjX3R5cGVzIHx8LS1veyBtZXJjX3RlbXBsYXRlcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | dbstring | [db_str](../../schema/client-files/db_str.md) | id |
| Has-Many | merc_type_id | [merc_templates](../../schema/mercenaries/merc_templates.md) | merc_type_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| merc_type_id | int | Unique Mercenary Type Identifier |
| race_id | int | [Race Identifier](../../../../categories/npc/race-list) |
| proficiency_id | tinyint | Proficiency Identifier |
| dbstring | varchar | [DBString Identifier](../../../schema/client-files/db_str.md) |
| clientversion | int | [Client Version](../../../../categories/player/client-version-bitmasks) |

