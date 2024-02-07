# merc_types

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY190eXBlcyB7XG4gICAgICAgIHZhcmNoYXIgZGJzdHJpbmdcbiAgICAgICAgdmFyY2hhciBtZXJjX3R5cGVfaWRcbiAgICB9XG4gICAgZGJfc3RyIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIG1lcmNfdGVtcGxhdGVzIHtcbiAgICAgICAgdmFyY2hhciBtZXJjX3N1YnR5cGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX3RlbXBsYXRlX2lkXG4gICAgICAgIHZhcmNoYXIgbWVyY190eXBlX2lkXG4gICAgICAgIHZhcmNoYXIgZGJzdHJpbmdcbiAgICAgICAgdmFyY2hhciBuYW1lX3R5cGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX25wY190eXBlX2lkXG4gICAgfVxuICAgIG1lcmNfdHlwZXMgfHwtLW97IGRiX3N0ciA6IFwiT25lLXRvLU9uZVwiXG4gICAgbWVyY190eXBlcyB8fC0tb3sgbWVyY190ZW1wbGF0ZXMgOiBcIkhhcy1NYW55XCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY190eXBlcyB7XG4gICAgICAgIHZhcmNoYXIgZGJzdHJpbmdcbiAgICAgICAgdmFyY2hhciBtZXJjX3R5cGVfaWRcbiAgICB9XG4gICAgZGJfc3RyIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIG1lcmNfdGVtcGxhdGVzIHtcbiAgICAgICAgdmFyY2hhciBtZXJjX3N1YnR5cGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX3RlbXBsYXRlX2lkXG4gICAgICAgIHZhcmNoYXIgbWVyY190eXBlX2lkXG4gICAgICAgIHZhcmNoYXIgZGJzdHJpbmdcbiAgICAgICAgdmFyY2hhciBuYW1lX3R5cGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX25wY190eXBlX2lkXG4gICAgfVxuICAgIG1lcmNfdHlwZXMgfHwtLW97IGRiX3N0ciA6IFwiT25lLXRvLU9uZVwiXG4gICAgbWVyY190eXBlcyB8fC0tb3sgbWVyY190ZW1wbGF0ZXMgOiBcIkhhcy1NYW55XCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY190eXBlcyB7XG4gICAgICAgIHZhcmNoYXIgZGJzdHJpbmdcbiAgICAgICAgdmFyY2hhciBtZXJjX3R5cGVfaWRcbiAgICB9XG4gICAgZGJfc3RyIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIG1lcmNfdGVtcGxhdGVzIHtcbiAgICAgICAgdmFyY2hhciBtZXJjX3N1YnR5cGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX3RlbXBsYXRlX2lkXG4gICAgICAgIHZhcmNoYXIgbWVyY190eXBlX2lkXG4gICAgICAgIHZhcmNoYXIgZGJzdHJpbmdcbiAgICAgICAgdmFyY2hhciBuYW1lX3R5cGVfaWRcbiAgICAgICAgdmFyY2hhciBtZXJjX25wY190eXBlX2lkXG4gICAgfVxuICAgIG1lcmNfdHlwZXMgfHwtLW97IGRiX3N0ciA6IFwiT25lLXRvLU9uZVwiXG4gICAgbWVyY190eXBlcyB8fC0tb3sgbWVyY190ZW1wbGF0ZXMgOiBcIkhhcy1NYW55XCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


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
| dbstring | varchar | [DBString Identifier](../../schema/client-files/db_str.md) |
| clientversion | int | [Client Version](../../../../categories/player/client-version-bitmasks) |

