# faction_list_mod

!!! info
	This page was last generated 2023.01.30

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZmFjdGlvbl9saXN0X21vZCB7XG4gICAgICAgIGludHVuc2lnbmVkIGZhY3Rpb25faWRcbiAgICB9XG4gICAgZmFjdGlvbl9saXN0IHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIGZhY3Rpb25fbGlzdF9tb2QgfHwtLW97IGZhY3Rpb25fbGlzdCA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZmFjdGlvbl9saXN0X21vZCB7XG4gICAgICAgIGludHVuc2lnbmVkIGZhY3Rpb25faWRcbiAgICB9XG4gICAgZmFjdGlvbl9saXN0IHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIGZhY3Rpb25fbGlzdF9tb2QgfHwtLW97IGZhY3Rpb25fbGlzdCA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZmFjdGlvbl9saXN0X21vZCB7XG4gICAgICAgIGludHVuc2lnbmVkIGZhY3Rpb25faWRcbiAgICB9XG4gICAgZmFjdGlvbl9saXN0IHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIGZhY3Rpb25fbGlzdF9tb2QgfHwtLW97IGZhY3Rpb25fbGlzdCA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | faction_id | [faction_list](../../schema/factions/faction_list.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Faction List Modifier Identifier |
| faction_id | int | [Faction Identifier](faction_list.md) |
| mod | smallint | Modifier |
| mod_name | varchar | Modifier Name: r# = [Race Identifier](../../../../server/npc/race-list), c# = [Class Identifier](../../../../server/player/class-list), d# = [Deity Identifier](../../../../server/player/deity-list) |

